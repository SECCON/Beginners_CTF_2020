package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-redis/redis"
	"github.com/rekkusu/c4b/app/contract"
)

type webServer struct {
	mux     *http.ServeMux
	manager *contractManager
	flag    string
	redis   *redis.Client
}

func NewServer(contractManager *contractManager, flag string, rclient *redis.Client) *webServer {
	server := &webServer{
		mux:     http.NewServeMux(),
		manager: contractManager,
		flag:    flag,
		redis:   rclient,
	}

	server.mux.HandleFunc("/", server.HandleIndex)
	server.mux.HandleFunc("/problem", server.HandleProblem)
	server.mux.HandleFunc("/get_challenge", server.HandleGetChallenge)
	server.mux.HandleFunc("/submit", server.HandleSubmit)

	return server
}

func (self *webServer) Start(listen string) error {
	return http.ListenAndServe(listen, self.mux)
}

func (self *webServer) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Signature string `json:"signature"`
		Address   string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	hash := accounts.TextHash([]byte("Did you enjoy this challenge?\n" + data.Address))
	sigBytes := common.FromHex(data.Signature)
	if sigBytes[len(sigBytes)-1] >= 27 {
		sigBytes[len(sigBytes)-1] = sigBytes[len(sigBytes)-1] - 27
	}
	pubkey, err := crypto.SigToPub(hash, sigBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	player := crypto.PubkeyToAddress(*pubkey)
	challengeAddr, err := self.redis.Get(strings.ToLower(player.String())).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = self.manager.CheckC4B(player, common.HexToAddress(challengeAddr))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"flag": self.flag,
	})
}

func (self *webServer) HandleGetChallenge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Player string `json:"player"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	challenge, err := self.redis.Get(strings.ToLower(data.Player)).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
	} else if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"challenge": challenge,
	})
}

func (self *webServer) HandleProblem(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("contracts/C4B.sol")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	problem := make(map[string]interface{})
	problem["deploy_address"] = self.manager.ChallengeManagerAddress.Hex()
	problem["contract_source"] = string(data)
	problem["abi"] = map[string]interface{}{
		"manager":   contract.ChallengeManagerABI,
		"challenge": contract.C4BABI,
	}

	resp, err := json.Marshal(problem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

func (self *webServer) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file := path.Join("web", r.URL.Path)
		http.ServeFile(w, r, file)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
