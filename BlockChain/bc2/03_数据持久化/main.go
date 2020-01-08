package main

import (
	"blockchain/03_数据持久化/BLC"
)

func main() {
	//创建创世区块
	blockchain := blc.CreateGenesisBlockChain("Genesis Block")
	defer blockchain.DB.Close()

	// 添加几个新区块
	blockchain.AddBlockToBlockChain("Send 100RMB to wangergou")
	blockchain.AddBlockToBlockChain("Send 100RMB to lixiaohua")
	blockchain.AddBlockToBlockChain("Send 100RMB to rose")

	// 递归打印区块数据
	blockchain.PrintChainV1()

}
