package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	// 引入以太坊客户端的库
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接到以太坊客户端节点，这里使用的是Cloudflare提供的以太坊节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		// 如果连接出错，输出错误信息并终止程序
		log.Fatal(err)
	}

	// 获取最新区块的头信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 打印最新区块的编号
	fmt.Println(header.Number.String()) // 输出示例：21138337

	// 创建一个特定区块号，这里指定区块号为 21138337
	blockNumber := big.NewInt(21138337)
	// 获取指定区块的信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 打印区块号
	fmt.Println(block.Number().Uint64()) // 示例输出：5671744
	// 打印区块的时间戳（Unix时间戳格式）
	fmt.Println(block.Time()) // 示例输出：1527211625
	// 打印区块的难度值
	fmt.Println(block.Difficulty().Uint64()) // 示例输出：3217000136609065
	// 打印区块的哈希值
	fmt.Println(block.Hash().Hex()) // 示例输出：0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	// 打印区块中的交易数量
	fmt.Println(len(block.Transactions())) // 示例输出：144

	// 根据区块哈希获取区块中的交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// 打印区块中的交易数量，和之前获取的交易数量应一致
	fmt.Println(count) // 示例输出：144
}
