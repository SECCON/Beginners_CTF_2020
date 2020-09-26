package main

import (
	"log"
	"os"

	"github.com/rekkusu/c4b/app"
)

//go:generate abigen --sol contracts/ChallengeManager.sol --pkg contract --out app/contract/challenge_manager.go

func main() {
	rpcurl, ok := os.LookupEnv("C4B_RPCURL")
	if !ok {
		rpcurl = "ws://localhost:7545"
	}

	redis, ok := os.LookupEnv("C4B_REDIS")
	if !ok {
		redis = "localhost:6379"
	}

	privkey, ok := os.LookupEnv("C4B_PRIVKEY")
	if !ok {
		privkey = "d00a997dc8924c035deeb2202675c4c26e3a7adcc662bf809389dcbe189b4f30"
	}

	flag, ok := os.LookupEnv("C4B_FLAG")
	if !ok {
		flag = "FLAG{sample}"
	}

	listen, ok := os.LookupEnv("C4B_LISTEN")
	if !ok {
		listen = "127.0.0.1:8000"
	}

	app := app.NewApp(app.AppConfig{
		RPC_URL:    rpcurl,
		Redis:      redis,
		PrivateKey: privkey,
		FLAG:       flag,
		Listen:     listen,
	})
	log.Fatal(app.Start())
}
