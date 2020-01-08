package BLC

import (
	"io"
	"bytes"
	"log"
	"net"
	"fmt"
)


//COMMAND_VERSION
func sendVersion(toAddress string,bc *BlockChain)  {


	bestHeight := bc.GetBestHeight()

	payload := gobEncode(Version{NODE_VERSION, bestHeight, nodeAddress})

	//version
	request := append(commandToBytes(COMMAND_VERSION), payload...)

	sendData(toAddress,request)


}



//COMMAND_GETBLOCKS
func sendGetBlocks(toAddress string)  {

	payload := gobEncode(GetBlocks{nodeAddress})

	request := append(commandToBytes(COMMAND_GETBLOCKS), payload...)

	fmt.Printf("toAddress:%s",toAddress)
	sendData(toAddress,request)

}

// 主节点将自己的所有的区块hash发送给钱包节点
//COMMAND_BLOCK
//
func sendInv(toAddress string, kind string, hashes [][]byte) {

	payload := gobEncode(Inv{nodeAddress,kind,hashes})

	request := append(commandToBytes(COMMAND_INV), payload...)

	sendData(toAddress,request)

}



func sendGetData(toAddress string, kind string ,blockHash []byte) {

	payload := gobEncode(GetData{nodeAddress,kind,blockHash})

	request := append(commandToBytes(COMMAND_GETDATA), payload...)

	sendData(toAddress,request)
}



func sendBlock(toAddress string, block []byte)  {


	payload := gobEncode(BlockData{nodeAddress,block})

	request := append(commandToBytes(COMMAND_BLOCK), payload...)

	sendData(toAddress,request)

}

func sendTx(toAddress string,tx *Transaction)  {


	payload := gobEncode(Tx{nodeAddress,tx})

	request := append(commandToBytes(COMMAND_TX), payload...)

	sendData(toAddress,request)

}

func sendData(to string,data []byte)  {

	conn, err := net.Dial(PROTOCOL, to)
	fmt.Println(err)
	if err != nil {
		panic("error")
	}
	defer conn.Close()

	// 附带要发送的数据
	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}

type Tx struct {
	AddrFrom string
	Tx *Transaction
}


//存储节点全局变量
//localhost:3000 主节点的地址
var knowNodes = []string{"localhost:3000"}
var nodeAddress string //全局变量，节点地址
// 存储hash值
var transactionArray [][]byte
var minerAddress string
var memoryTxPool = make(map[string]*Transaction)


type Version struct {
	Version    int64 // 版本
	BestHeight int64 // 当前节点区块的高度
	AddrFrom   string //当前节点的地址
}