package blc

import (
	"blockchain/common"
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//step1:创建Block结构体
type Block struct {
	//高度Height：其实就是区块的编号，第一个区块叫创世区块，高度为0
	Height int64
	//上一个区块的哈希值ProvHash：
	PrevBlockHash []byte
	//交易数据Data：目前先设计为[]byte,后期是Transaction
	Data [] byte
	//时间戳TimeStamp：
	TimeStamp int64
	//哈希值Hash：32个的字节，64个16进制数
	Hash []byte
}

//step2：创建新的区块
func NewBlock(data string, provBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, provBlockHash, []byte(data), time.Now().Unix(), nil}
	//设置哈希值
	block.setHash()
	return block
}

//step3:设置区块的hash
func (block *Block) setHash() {
	//1.将高度转为字节数组
	heightBytes := common.IntToHex(block.Height)

	//2.时间戳转为字节数组
	timeString := strconv.FormatInt(block.TimeStamp, 2)

	timeBytes := [] byte(timeString)
	//fmt.Println("timeStamp:",timeBytes)

	//3.拼接所有的属性
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes}, []byte{})

	//4.生成哈希值
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

//step4:创建创世区块：
func GenesisBlock() *Block {
	return NewBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", make([] byte, 32, 32), 0)
}

