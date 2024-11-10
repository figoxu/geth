package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	account := common.HexToAddress("0x47B3b6a0E19e908e0b4500f63BE4c1E6d46C52d4")

	fmt.Println(account.Hex())   // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(account.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 查看账户余额

	pendingBalanceAt, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		return
	}
	fmt.Println(pendingBalanceAt) // 查看账户余额
}
