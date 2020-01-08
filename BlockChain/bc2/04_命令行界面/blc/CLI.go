package blc

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct{}

func (cli *Cli) Run() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	//1. 创建【二级】命令参数【集合】
	create := flag.NewFlagSet("create", flag.ExitOnError) //创建
	add := flag.NewFlagSet("add", flag.ExitOnError)       //添加
	print := flag.NewFlagSet("print", flag.ExitOnError)   //打印

	flagAdd := add.String("d", "block data", "交易数据")
	flagCreate := create.String("d", "Genesis data", "创世区块")

	//2.解析
	switch os.Args[1] {
	case "add":
		err := add.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
		//fmt.Println("----", os.Args[2:])
	case "print":
		err := print.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "create":
		err := create.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1) //退出
	}

	// 3.解析成功
	if add.Parsed() {
		if *flagAdd == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAdd)
	}
	if print.Parsed() {
		cli.printChains()
	}

	if create.Parsed() {
		if *flagCreate == "" {
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreate)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreate -d DATA -- 创建创世区块")
	fmt.Println("\tadd -d Data -- 交易数据")
	fmt.Println("\tprint -- 输出信息")
}

func (cli *Cli) addBlock(data string) {
	bc := GetBlockchainObject()
	if bc == nil {
		fmt.Println("没有创世区块，无法添加。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.AddBlockToBlockChain(data)
}

func (cli *Cli) printChains() {
	bc := GetBlockchainObject()
	if bc == nil {
		fmt.Println("没有区块可以打印。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.PrintChains()
}

func (cli *Cli) createGenesisBlockchain(data string) {
	CreateBlockChainWithGenesisBlock(data)
}
