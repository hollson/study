package BLC

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

// 对blockchain的命令行操作进行管理

// client对象
type CLI struct {}

// 用法展示
func PrintUsage() {
	fmt.Println("Usage:")
	// 初始化区块链
	fmt.Printf("\tcreateblockchain -address address -- 创建区块链\n")
	// 添加区块
	//fmt.Printf("\taddblock -data DATA -- 添加区块\n")
	// 打印完整的区块信息
	fmt.Printf("\tprintchain -- 输出区块链信息\n")
	// 通过命令转账
	fmt.Printf("\tsend -from FROM -to TO -amount AMOUNT -- 发起转账\n")
	fmt.Printf("\t转账参数说明\n")
	fmt.Printf("\t\t-from FROM -- 转账源地址\n")
	fmt.Printf("\t\t-to TO -- 转账目标地址\n")
	fmt.Printf("\t\t-AMOUNT amount -- 转账金额\n")
	// 查询余额
	fmt.Printf("\tgetbalance -address FROM -- 查询指定地址的余额\n")
	fmt.Println("\t查询余额参数说明")
	fmt.Printf("\t\t-address -- 查询余额的地址")
}

// 添加区块
func (cli *CLI) addBlock(txs []*Transaction) {
	// 判断数据库是否存在
	if !dbExist() {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	// 获取到blockchain的对象实例
	blockchain.AddBlock(txs)
}

// 命令行运行函数
func (cli *CLI) Run() {
	// 检测参数数量
	IsValidArgs()
	// 新建相关命令
	// 添加区块
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	// 输出区块链完整信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	// 创建区块链
	createBLCWithGenesisBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	// 发起交易
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	// 查询余额命令
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	// 数据参数处理
	// 添加区块参数
	flagAddBlockArg := addBlockCmd.String("data", "sent 100 btc to player", "添加区块数据" )
	// 创建区块链时指定的矿工地址(接收奖励)参数
	flagCreateBlockchainArg := createBLCWithGenesisBlockCmd.String("address",
		"troytan","指定接收系统奖励的矿工地址")
	// 发起交易参数
	flagSendFromArg := sendCmd.String("from", "", "转账源地址")
	flagSendToArg := sendCmd.String("to", "", "转账目标地址")
	flagSendAmountArg := sendCmd.String("amount", "", "转账金额")
	// 查询余额命令行参数
	flagGetBalanceArg := getBalanceCmd.String("address", "", "要查询的地址")
	// 判断命令
	switch os.Args[1] {
	case "getbalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse cmd get balance failed! %v\n", err)
		}
	case "send":
		if err := sendCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse sendCmd failed! %v\n", err)
		}
	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse addBlockCmd failed! %v\n", err)
		}
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse printchainCmd failed! %v\n", err)
		}
	case "createblockchain":
		if err := createBLCWithGenesisBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse createBLCWithGenesisBlockCmd failed! %v\n", err)
		}
	default:
		// 没有传递任何命令或者传递的命令不在上面的命令列表之中
		PrintUsage()
		os.Exit(1)
		
	}
	// 查询余额
	if getBalanceCmd.Parsed() {
		if *flagGetBalanceArg == "" {
			fmt.Println("请输入查询地址...")
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceArg)
	}
	// 发起转账
	if sendCmd.Parsed() {
		if *flagSendFromArg == "" {
			fmt.Println("源地址不能为空...")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendToArg == "" {
			fmt.Println("目标地址不能为空...")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendAmountArg == "" {
			fmt.Println("转账金额不能为空...")
			PrintUsage()
			os.Exit(1)
		}
		fmt.Printf("\tFROM:[%s]\n", JSONToSlice(*flagSendFromArg))
		fmt.Printf("\tTO:[%s]\n", JSONToSlice(*flagSendToArg))
		fmt.Printf("\tAMOUNT:[%s]\n", JSONToSlice(*flagSendAmountArg))
		cli.send(JSONToSlice(*flagSendFromArg), JSONToSlice(*flagSendToArg), JSONToSlice(*flagSendAmountArg))
	}
	// 添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		// 调用
		cli.addBlock([]*Transaction{})
	}
	// 输出区块链信息
	if printChainCmd.Parsed() {
		cli.printchain()
	}
	// 创建区块链命令
	if createBLCWithGenesisBlockCmd.Parsed() {
		if *flagCreateBlockchainArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockchain(*flagCreateBlockchainArg)
	}
}