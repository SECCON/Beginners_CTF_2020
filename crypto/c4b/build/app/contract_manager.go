package app

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rekkusu/c4b/app/contract"
)

type contractManager struct {
	client                  *ethclient.Client
	privKey                 *ecdsa.PrivateKey
	pubKey                  *ecdsa.PublicKey
	ChallengeManager        *contract.ChallengeManager
	ChallengeManagerAddress common.Address
}

func ConnectNetwork(RPC_URL string, keyString string) (*contractManager, error) {
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.HexToECDSA(keyString)
	if err != nil {
		return nil, err
	}

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if ok == false {
		return nil, errors.New("key error")
	}

	go func() {
		for {
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("PING: Current block is %s\n", header.Number.String())
			time.Sleep(1 * time.Hour)
		}
	}()

	return &contractManager{
		client:  client,
		pubKey:  pubKeyECDSA,
		privKey: privKey,
	}, nil
}

func (self *contractManager) InitContract() (string, error) {
	fromAddress := crypto.PubkeyToAddress(*self.pubKey)

	nonce, err := self.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := self.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(self.privKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1000000)
	auth.GasPrice = gasPrice

	log.Println("Deploying ChallengeManager")
	address, tx, instance, err := contract.DeployChallengeManager(auth, self.client)
	if err != nil {
		return "", err
	}
	log.Printf("Deployed ChallengeManager at %s via %s\n", address.Hex(), tx.Hash().Hex())

	self.ChallengeManager = instance
	self.ChallengeManagerAddress = address

	return address.Hex(), nil
}

func (self *contractManager) LoadContract(address string) error {
	contractAddr := common.HexToAddress(address)
	instance, err := contract.NewChallengeManager(contractAddr, self.client)
	if err != nil {
		return err
	}

	log.Printf("Loaded ChallengeManager at %s\n", contractAddr.Hex())
	self.ChallengeManager = instance
	self.ChallengeManagerAddress = contractAddr
	return nil
}

func (self *contractManager) CheckC4B(player common.Address, challenge common.Address) error {
	instance, err := contract.NewC4B(challenge, self.client)
	if err != nil {
		log.Print(err)
		return errors.New("Unknown error")
	}

	p, err := instance.Player(nil)
	if err != nil {
		log.Print(err)
		return errors.New("Unknown error")
	}

	if !bytes.Equal(p.Bytes(), player.Bytes()) {
		return errors.New("this contract is not yours")
	}

	success, err := instance.Success(nil)
	if err != nil {
		log.Print(err)
		return errors.New("Unknown error")
	}

	if !success {
		return errors.New("this contract does not satisfy the condition")
	}

	return nil
}

func (self *contractManager) WatchChallengeDeployed() (<-chan *contract.ChallengeManagerChallengeDeployed, event.Subscription, error) {
	ch := make(chan *contract.ChallengeManagerChallengeDeployed)
	subs, err := self.ChallengeManager.WatchChallengeDeployed(&bind.WatchOpts{}, ch, nil)
	if err != nil {
		return nil, nil, err
	}

	return ch, subs, nil
}
