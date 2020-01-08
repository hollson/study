package main

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main()  {
	// sha256
	hash := sha256.New()
	hash.Write([]byte("eth1804"))
	bytes := hash.Sum(nil)
	fmt.Printf("sha256 : %x\n", bytes)

	// 3eaba0d67ee9c856c26fab143b9c61bb2884af68579898054a61a6ece1fbe6dd
	// 3eaba0d67ee9c856c26fab143b9c61bb2884af68579898054a61a6ece1fbe6dd

	// ripemd160
	r160 := ripemd160.New()
	r160.Write(bytes)
	bytesRipemd := r160.Sum(nil)
	fmt.Printf("rpiemd160 : %x\n", bytesRipemd)
}
