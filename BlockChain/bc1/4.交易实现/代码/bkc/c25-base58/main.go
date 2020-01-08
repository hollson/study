package main

import (
	"eth-1804/bkc/c25-base58/BLC"
	"fmt"
)

func main() {
	result := BLC.Base58Encode([]byte("this is the example"))
	fmt.Printf("result : %s\n", result)
}
