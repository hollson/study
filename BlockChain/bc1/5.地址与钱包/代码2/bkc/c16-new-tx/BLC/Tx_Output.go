package BLC

// 交易的输出管理

// 输出结构
type TxOutput struct {
	// 金额(大写才能导出金额)
	Value				int
	// 用户名(UTXO的所有者)
	ScriptPubkey		string
}