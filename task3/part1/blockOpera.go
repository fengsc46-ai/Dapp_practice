package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	queryBlock1()
}

func queryBlock1() {
	// 连接Solana Testnet RPC
	client := rpc.New(rpc.TestNet_RPC)

	ctx := context.Background()

	// 1️⃣ 获取最新 Slot
	latestSlot, err := client.GetSlot(ctx, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("获取 Slot 失败: %v", err)
	}
	fmt.Printf("最新 Slot: %d\n", latestSlot)

	// 2️⃣ 根据 Slot 获取区块信息
	blockResp, err := client.GetBlock(ctx, latestSlot)
	if err != nil {
		log.Fatalf("获取区块信息失败: %v", err)
	}

	// 输出基本信息
	fmt.Printf("区块哈希: %s\n", blockResp.Blockhash)
	fmt.Printf("上一个区块哈希: %s\n", blockResp.PreviousBlockhash)
	fmt.Printf("区块时间: %v\n", blockResp.BlockTime)
}
