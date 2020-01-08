package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/labstack/gommon/log"
	"time"
)

// 区块基本结构与功能管理文件
// 实现一个最基本的区块结构
type Block struct {
	TimeStamp		int64		//区块时间戳，代表区块时间
	Hash			[]byte		// 当前区块哈希
	PrevBlockHash	[]byte		// 前区块哈希
	Heigth			int64		// 区块高度
	Data			[]byte		// 交易数据
	Nonce			int64		// 在运行pow时生成的哈希变化值，也代表pow运行时动态修改的数据
}

// 新建区块
func NewBlock(height int64, prevBlockHash []byte, data []byte) *Block {
	var block Block
	block = Block{
		TimeStamp:time.Now().Unix(),
		Hash: nil,
		PrevBlockHash:prevBlockHash,
		Heigth:height,
		Data:data,
	}
	// 生成哈希
	block.SetHash()
	// 替换setHash
	// 通过POW生成新的哈希
	pow := NewProofOfWork(&block)
	// 执行工作量证明算法
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	return &block
}

// 计算区块哈希
func (b *Block) SetHash()  {
	// 调用sha256实现哈希生成
	// 实现int64->hash
	timeStampBytes := IntToHex(b.TimeStamp)
	heigthBytes := IntToHex(b.Heigth)
	blockBytes := bytes.Join([][]byte{
		heigthBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	},[]byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}

// 生成创世区块
func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}

// 区块结构序列化
func (block *Block)Serialize() []byte {
	var buffer bytes.Buffer
	// 新建编码对象
	encoder := gob.NewEncoder(&buffer)
	// 编码(序列化)
	if err := encoder.Encode(block); nil != err {
		log.Panicf("serialize the block to []byte failed! %v\n", err)
	}
	return buffer.Bytes()
}
// 区块数据返序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	// 新建 decoder对象
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); nil != err {
		log.Panicf("deserialize the []byte to block failed! %v\n", err)
	}
	return &block
}