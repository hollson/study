package main

import (
	"eth-1804/bkc/c26-wallet/BLC"
	"fmt"
)

func main() {
	result := BLC.Base58Encode([]byte("this is the example"))
	fmt.Printf("result : %s\n", result)
	decodeResult := BLC.Base58Decode([]byte("1nj2SLMErZakmBni8xhSXtimREn"))
	fmt.Printf("decode result : %s\n", decodeResult)
}
