package main

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

func main() {
	// 定义助记词，用于生成钱包的主密钥
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"

	// 根据助记词创建HD钱包实例
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err) // 如果创建钱包失败，则记录错误并退出
	}

	// 设置第一个派生路径，表示第一个地址
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	// 从钱包中根据派生路径生成第一个账户地址
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err) // 如果地址生成失败，则记录错误并退出
	}

	// 打印生成的第一个账户地址
	fmt.Println(account.Address.Hex()) // 输出：0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	// 设置第二个派生路径，表示第二个地址
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	// 从钱包中根据派生路径生成第二个账户地址
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err) // 如果地址生成失败，则记录错误并退出
	}

	// 打印生成的第二个账户地址
	fmt.Println(account.Address.Hex()) // 输出：0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}
