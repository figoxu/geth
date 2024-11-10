package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NewKeyStore：用于指定存储路径，示例中使用了/tmp/keystore。
// NewAccount：创建新账户，使用密码加密。
// Export：导出加密的私钥。
// DecryptKey：解密密钥文件，从而获取私钥。
func main() {
	// 创建一个新的临时 keystore（可以更改路径将其保存在本地）
	ks := keystore.NewKeyStore("/Users/zhangxiazhen/figoxu/leet_code/web3/geth/resources/keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// 设置账户的密码
	password := "xujianhui@bj"

	// 创建账户
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalf("无法创建账户: %v", err)
	}
	fmt.Println("新账户地址:", account.Address.Hex())

	// 解锁账户并获取私钥
	key, err := ks.Export(account, password, password)
	if err != nil {
		log.Fatalf("无法导出私钥: %v", err)
	}

	// 使用 keystore 的 DecodeKey 解析私钥
	privKey, err := keystore.DecryptKey(key, password)
	if err != nil {
		log.Fatalf("无法解密私钥: %v", err)
	}

	privateKeyECDSA := privKey.PrivateKey
	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)

	fmt.Println("私钥:", hexutil.Encode(privateKeyBytes))
}
