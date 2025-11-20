package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"practice/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// import (
//
//	"context"
//	"fmt"
//	"log"
//
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/ethclient"
//
// )
//
//	func main() {
//		client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/36VCcRTF-M8cAMJFumpHi")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		headers := make(chan *types.Header)
//		sub, err := client.SubscribeNewHead(context.Background(), headers)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for {
//			select {
//			case err := <-sub.Err():
//				log.Fatal(err)
//			case header := <-headers:
//				fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
//				block, err := client.BlockByHash(context.Background(), header.Hash())
//				if err != nil {
//					log.Fatal(err)
//				}
//
//				fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
//				fmt.Println(block.Number().Uint64())   // 3477413
//				fmt.Println(block.Time())              // 1529525947
//				fmt.Println(block.Nonce())             // 130524141876765836
//				fmt.Println(len(block.Transactions())) // 7
//			}
//		}
//	}
func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/36VCcRTF-M8cAMJFumpHi")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.GenerateKey()
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
	// fmt.Println("Private Key:", privateKeyHex)
	privateKey, err := crypto.HexToECDSA("<private_key>")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := token.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
