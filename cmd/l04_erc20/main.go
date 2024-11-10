package main

import (
	"fmt"
	token "geth/cmd/l04_erc20/contracts_erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

// 2.Go以太坊开发/05.账户代币余额/账户代币余额.md:10
func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误
	}

	// 定义 Golem (GNT) 的合约地址
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")

	// 创建一个代币合约实例，用于调用合约中的方法
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err) // 如果创建实例失败，记录错误
	}

	// 定义要查询余额的地址
	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")

	// 调用合约的 BalanceOf 方法查询该地址的余额
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err) // 如果查询失败，记录错误
	}

	// 调用合约的 Name 方法获取代币的名称
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err) // 如果查询失败，记录错误
	}

	// 调用合约的 Symbol 方法获取代币的符号
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err) // 如果查询失败，记录错误
	}

	// 调用合约的 Decimals 方法获取代币的小数位数
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err) // 如果查询失败，记录错误
	}

	// 打印代币的基本信息
	fmt.Printf("name: %s\n", name)         // 代币名称，例如 "Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // 代币符号，例如 "GNT"
	fmt.Printf("decimals: %v\n", decimals) // 代币小数位数，例如 18

	// 打印地址的代币余额，单位为 wei
	fmt.Printf("wei: %s\n", bal) // 余额的原始 wei 值，例如 "74605500647408739782407023"

	// 将余额转换为 float 类型，并除以 10 的 decimals 次方，得到代币单位下的余额
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	// 打印最终的余额，以代币单位显示，例如 "balance: 74605500.647409"
	fmt.Printf("balance: %f", value)
}
