package main

import "eth-1804/bkc/c13-cli-send/BLC"

// 启动
func main() {
	//bc := BLC.CreateBlockChainWithGenesisBlock()
	////fmt.Printf("blockchain:%v\n", bc.Blocks[0])
	//bc.AddBlock([]byte("a send 100 eth to b"))
	//
	//bc.PrintChain()
	cli := BLC.CLI{}
	cli.Run()
}