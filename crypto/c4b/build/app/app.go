package app

import (
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/go-redis/redis"
)

const (
	ChallengeManagerPath = "challenge_manager.txt"
)

type app struct {
	config AppConfig
	Redis  *redis.Client
}

type AppConfig struct {
	RPC_URL    string
	PrivateKey string
	StateDir   string
	FLAG       string
	Listen     string
	Redis      string
}

func NewApp(config AppConfig) *app {
	return &app{
		config: config,
	}
}

func (self *app) Start() error {
	log.Println("Redis Setup")

	self.Redis = redis.NewClient(&redis.Options{
		Addr: self.config.Redis,
	})
	_, err := self.Redis.Ping().Result()
	if err != nil {
		return err
	}

	log.Println("Connecting Ethereum API")
	manager, err := ConnectNetwork(self.config.RPC_URL, self.config.PrivateKey)
	if err != nil {
		return err
	}

	err = initContractIfNotExists(manager, self.config.StateDir)
	if err != nil {
		return err
	}

	err = handleChallengeDeployed(manager, self.Redis)
	if err != nil {
		return err
	}

	log.Println("Launching Server")
	server := NewServer(manager, self.config.FLAG, self.Redis)
	return server.Start(self.config.Listen)
}

func initContractIfNotExists(manager *contractManager, stateDir string) error {
	data, err := ioutil.ReadFile(path.Join(stateDir, ChallengeManagerPath))
	if err == nil && len(data) > 10 {
		log.Printf("Load Contract at %s\n", string(data))
		return manager.LoadContract(string(data))
	}

	// File not found
	address, err := manager.InitContract()
	if err != nil {
		return err
	}
	ioutil.WriteFile(path.Join(stateDir, ChallengeManagerPath), []byte(address), 0644)

	return nil
}

func handleChallengeDeployed(manager *contractManager, rclient *redis.Client) error {
	evCh, subs, err := manager.WatchChallengeDeployed()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-evCh:
				log.Printf("New Challenge was deployed: %s\n", event.Challenge.String())
				err := rclient.Set(strings.ToLower(event.Player.String()), event.Challenge.String(), 0).Err()
				if err != nil {
					log.Fatal(err)
				}
			case err := <-subs.Err():
				log.Fatal(err)
			}
		}
	}()

	return nil
}
