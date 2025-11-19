package Dapp_practice

//func main() {
//	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/36VCcRTF-M8cAMJFumpHi")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	value := big.NewInt(1000000000000000000) // in wei (1 eth)
//	gasLimit := uint64(21000)                // in units
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	toAddress := common.HexToAddress("0x5d64e34bc67c2cc73da25b7b98f959811bea5851")
//	var data []byte
//
//	tx := types.NewTx(&types.LegacyTx{
//		Nonce:    nonce,
//		GasPrice: gasPrice,
//		Gas:      gasLimit,
//		To:       &toAddress,
//		Value:    value,
//		Data:     data,
//	})
//
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.SendTransaction(context.Background(), signedTx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
//}

// 账户余额查询
//func main() {
//	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/36VCcRTF-M8cAMJFumpHi")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	account := common.HexToAddress("0x6009be5bd6c1b3091fe465134a5ee270f23362d8")
//	balance, err := client.BalanceAt(context.Background(), account, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(balance)
//	blockNumber := big.NewInt(5532993)
//	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(balanceAt) // 25729324269165216042
//	fbalance := new(big.Float)
//	fbalance.SetString(balanceAt.String())
//	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
//	fmt.Println(ethValue) // 25.729324269165216041
//	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
//	fmt.Println(pendingBalance) // 25729324269165216042
//}
