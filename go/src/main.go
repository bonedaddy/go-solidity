package main


import (
	"log"
	"strings"
	"fmt"
	"time"

	//ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/ethclient"
	"./eth_events"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

func main() {
	client, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatalf("error connecting to clent instance")
	}

	// authorize a conenction so we can deploy stuff and change states
	auth, err := bind.NewTransactor(strings.NewReader(key), "password123")
	if err != nil {
		log.Fatalf("error unlocking account")
	}

	address, tx, event_emitter, err := eth_events.DeployEthEvents(auth, client)
	if err != nil {
		log.Fatalf("error deploying event test")
	}
	fmt.Printf("Contract Address: 0x%x\nTransaction Hash: 0x%x\n", address, tx.Hash())

	fmt.Printf("sleeping for 1 minute")
	time.Sleep(1 * time.Minute)

	fmt.Printf("Emitting event")
	tx, err = event_emitter.EmitEvent(auth)
	if err != nil {
		log.Fatalf("error")
	}
	fmt.Printf("Transaction hash: 0x%x\n", tx.Hash())


	fmt.Printf("Loading event ABI so we can parse event logs\n")
    tokenAbi, err := abi.JSON(strings.NewReader(eth_events.EthEventsABI))
    if err != nil {
    	log.Fatalf("Error reading abi file")
    }
	fmt.Printf("ABI %s\n", tokenAbi)


	var ch = make(chan *eth_events.EthEventsTestEvent)

	// returns an event subscription
	sub, err := event_emitter.WatchTestEvent(nil, ch)
	if err != nil {
		log.Fatalf("error")
	}

	go eventParser(sub, ch)
	fmt.Println("event parser goroutine launched")
	fmt.Println("Going to sleep for 1 minute")
	time.Sleep(1 * time.Minute)
    /*for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case log := <-ch:
        	// the following prints the raw log data
            // fmt.Println("Log:", log)
            // the following prints log data for a
            fmt.Printf("unpacked log data A %s\n", log.A)
        }
    }*/	
}


func eventParser(sub event.Subscription, c chan *eth_events.EthEventsTestEvent) {
	
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-c:
			fmt.Printf("Unpacked log data A %s\nUnpacked log data B %s\n", log.A, log.B)
		}
	}
}