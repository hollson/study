package main

import "eth-1804/bkc/c17-cli-utxo/BLC"

// 启动
func main() {
	cli := BLC.CLI{}
	cli.Run()
}