package main

import "eth-1804/bkc/c18-get-utxo/BLC"

// 启动
func main() {
	cli := BLC.CLI{}
	cli.Run()
}