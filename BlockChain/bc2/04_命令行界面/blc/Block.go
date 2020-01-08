package blc

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Height        int64   //高度
	PrevBlockHash []byte  //前一区块
	Data          [] byte //交易数据
	TimeStamp     int64   //时间戳
	Hash          []byte  //哈希值
	Nonce         int64   //随机数
}

func NewBlock(data string, provBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, provBlockHash, []byte(data), time.Now().Unix(), nil, 0}
	//step5：设置block的hash和nonce
	//设置哈希
	//block.SetHash()
	//调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([] byte, 32, 32), 0)
}

// 序列化
func (block *Block) Serialize() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf) //gob编码器
	err := encoder.Encode(block)    //将对象编码打包
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

//反序列化
func DeserializeBlock(blockBytes [] byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	decoder := gob.NewDecoder(reader) //gob解码器
	err := decoder.Decode(&block)     //对象解包
	if err != nil {
		log.Panic(err)
	}
	return &block
}
