package main

import (
	"eth-1804/bkc/c06-printchain/BLC"
)

// 启动
func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()
	//fmt.Printf("blockchain:%v\n", bc.Blocks[0])
	bc.AddBlock([]byte("a send 100 eth to b"))

	bc.PrintChain()
	//bc.AddBlock([]byte("b send 100 eth to b"))
	// 上链
	// AddBlock(height int64, prevBlockHash []byte, data []byte)
	//bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Heigth + 1,
	//	bc.Blocks[len(bc.Blocks) - 1].Hash, []byte("alice send 10 btc to bob"))
	//bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Heigth + 1,
	//	bc.Blocks[len(bc.Blocks) - 1].Hash, []byte("bob send 5 to troytan"))
	//
	//for _, block := range bc.Blocks {
	//	fmt.Printf("prevBlockHash : %x, currentHash : %x\n",
	//		block.PrevBlockHash,  block.Hash)
	//}
}