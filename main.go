package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func main() {
	addressPrefix := "carauction"

	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
	)

	if err != nil {
		log.Fatal(err)
	}
	
	
	for {
		err := cosmos.WaitForNextBlock(context.Background())
		if err != nil {
			continue
		}

		blockNumber, err := cosmos.LatestBlockHeight(context.Background())
		if err != nil {
			continue
		}
		fmt.Println("Block number: ", blockNumber)
		
		txs, err := cosmos.GetBlockTXs(context.Background(), blockNumber)
		if err != nil {
			continue
		}
		
		for _, tx := range txs {
			events, _ := tx.GetEvents()
			fmt.Println("Events: ", events)
		}
	}
}