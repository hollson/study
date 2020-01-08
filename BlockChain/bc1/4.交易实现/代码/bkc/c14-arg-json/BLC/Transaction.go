package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/labstack/gommon/log"
)

// 交易管理文件

// 定义一个交易基本结构
type Transaction struct {
	// 交易哈希(标识)
	TxHash			[]byte
	// 输入列表
	Vins			[]*TxInput
	// 输出列表
	Vouts			[]*TxOutput
}

// 实现coinbase交易
func NewCoinbaseTransaction(address string) *Transaction {

	// 输入
	// coinbase特点
	// txHash:nil
	// vout: -1(为了对是否是coinbase交易进行判断)
	// ScriptSig:系统奖励
	txInput := &TxInput{[]byte{}, -1, "system reward"}
	// 输出：
	// value:
	// address
	txOutput := &TxOutput{10, address}
	// 输入输出组装交易
	txCoinbase := &Transaction{
		nil,
	[]*TxInput{txInput},
	[]*TxOutput{txOutput},
	}
	// 交易哈希生成
	txCoinbase.HashTransaction()
	return txCoinbase
}

// 生成交易哈希(交易序列化)
func (tx *Transaction) HashTransaction()  {
	var result bytes.Buffer
	// 设置编码对象
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(tx); err != nil {
		log.Panicf("tx Hash encoded failed %v\n", err)
	}

	// 生成哈希值
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}