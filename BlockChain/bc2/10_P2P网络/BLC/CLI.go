package BLC

import (
	"os"
	"fmt"
	"flag"
	"log"
	"strconv"
)

//step1:
//CLI结构体
type CLI struct {
	//Blockchain *BlockChain
}

//step2：添加Run方法
func (cli *CLI) Run() {
	//判断命令行参数的长度
	isValidArgs()

	/*
	获取节点ID
	解释：返回当前进程的环境变量varname的值,若变量没有定义时返回nil
	export NODE_ID=8888

	每次打开一个终端，都需要设置NODE_ID的值。
	变量名NODE_ID，可以更改别的。
	 */

	nodeID :=os.Getenv("NODE_ID")
	if nodeID == ""{
		fmt.Printf("NODE_ID 环境变量没有设置。。\n")
		os.Exit(1)
	}
	fmt.Printf("NODE_ID:%s\n",nodeID)


	//1.创建flagset标签对象
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	addressListsCmd := flag.NewFlagSet("addresslists",flag.ExitOnError)

	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	//fmt.Printf("%T\n",addBlockCmd) //*FlagSet
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	testCmd:=flag.NewFlagSet("test",flag.ExitOnError)

	startNodeCmd := flag.NewFlagSet("startnode",flag.ExitOnError)

	//2.设置标签后的参数
	//flagAddBlockData:= addBlockCmd.String("data","helloworld..","交易数据")
	flagFromData := sendBlockCmd.String("from", "", "转帐源地址")
	flagToData := sendBlockCmd.String("to", "", "转帐目标地址")
	flagAmountData := sendBlockCmd.String("amount", "", "转帐金额")
	flagCreateBlockChainData := createBlockChainCmd.String("address", "", "创世区块交易地址")
	flagGetBalanceData := getBalanceCmd.String("address", "", "要查询的某个账户的余额")



	flagMiner := startNodeCmd.String("miner","","定义挖矿奖励的地址......")
	flagMine := sendBlockCmd.Bool("mine",false,"是否在当前节点中立即验证....")


	//3.解析
	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
		//fmt.Println("----",os.Args[2:])

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
		//fmt.Println("====",os.Args[2:])

	case "createblockchain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addresslists":
		err := addressListsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "test":
		err := testCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1) //退出
	}

	if sendBlockCmd.Parsed() {
		if *flagFromData == "" || *flagToData == "" || *flagAmountData == "" {
			printUsage()
			os.Exit(1)
		}
		//cli.addBlock([]*Transaction{})
		//fmt.Println(*flagFromData)
		//fmt.Println(*flagToData)
		//fmt.Println(*flagAmountData)
		//fmt.Println(JSONToArray(*flagFrom))
		//fmt.Println(JSONToArray(*flagTo))
		//fmt.Println(JSONToArray(*flagAmount))
		from := JSONToArray(*flagFromData)
		to := JSONToArray(*flagToData)
		amount := JSONToArray(*flagAmountData)

		for i := 0; i < len(from); i++ {
			if !IsValidForAddress([]byte(from[i])) || !IsValidForAddress([]byte(to[i])) {
				fmt.Println("钱包地址无效")
				printUsage()
				os.Exit(1)
			}
		}

		cli.send(from, to, amount,nodeID,*flagMine)
	}
	if printChainCmd.Parsed() {
		cli.printChains(nodeID)
	}

	if createBlockChainCmd.Parsed() {
		//if *flagCreateBlockChainData == "" {
		if !IsValidForAddress([]byte(*flagCreateBlockChainData)){
			fmt.Println("创建地址无效")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainData,nodeID)
	}

	if getBalanceCmd.Parsed() {
		//if *flagGetBalanceData == "" {
		if !IsValidForAddress([]byte(*flagGetBalanceData)){
			fmt.Println("查询地址无效")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceData,nodeID)

	}

	if createWalletCmd.Parsed() {
		//创建钱包
		cli.createWallet(nodeID)
	}

	//获取所有的钱包地址
	if addressListsCmd.Parsed(){
		cli.addressLists(nodeID)
	}

	if testCmd.Parsed(){
		cli.TestMethod(nodeID)
	}

	if startNodeCmd.Parsed() {
		cli.startNode(nodeID,*flagMiner)
	}

}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreatewallet -- 创建钱包")
	fmt.Println("\taddresslists -- 输出所有钱包地址")
	fmt.Println("\tcreateblockchain -address DATA -- 创建创世区块")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -mine -- 交易明细.")
	fmt.Println("\tprintchain - 输出信息:")
	fmt.Println("\tgetbalance -address DATA -- 查询账户余额")
	fmt.Println("\ttest -- 测试")
	fmt.Println("\tstartnode -miner ADDRESS -- 启动节点服务器，并且指定挖矿奖励的地址.")
}


// 转账
func (cli *CLI) send(from []string, to []string, amount []string, nodeID string, mineNow bool) {

	blockchain := GetBlockchainObject(nodeID)
	utxoSet := &UTXOSet{blockchain}
	defer blockchain.DB.Close()

	if mineNow {
		blockchain.MineNewBlock(from, to, amount, nodeID)

		//转账成功以后，需要更新一下
		utxoSet.Update()
	} else {
		// 把交易发送到矿工节点去进行验证
		fmt.Println("由矿工节点处理......")
		value, _ := strconv.Atoi(amount[0])
		tx := NewSimpleTransaction(from[0], to[0], int64(value), utxoSet, []*Transaction{}, nodeID)

		sendTx(knowNodes[0], tx)
	}

}

func (cli *CLI) TestMethod(nodeID string){
	blockchain:=GetBlockchainObject(nodeID)
	//defer blockchain.DB.Close()

	unSpentOutputMap:=blockchain.FindUnSpentOutputMap()
	fmt.Println(unSpentOutputMap)
	for key,value:=range unSpentOutputMap{
		fmt.Println(key)
		for _,utxo:=range value.UTXOS{
			fmt.Println("金额：",utxo.Output.Value)
			fmt.Printf("地址：%v\n",utxo.Output.PubKeyHash)
			fmt.Println("---------------------")
		}
	}

	utxoSet:=&UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}

func (cli *CLI) startNode(nodeID string,minerAdd string)  {

	// 启动服务器
	fmt.Println(nodeID,minerAdd)

	if minerAdd == "" || IsValidForAddress([]byte(minerAdd))  {
		//  启动服务器
		fmt.Printf("启动服务器:localhost:%s\n",nodeID)
		startServer(nodeID,minerAdd)

	} else {

		fmt.Println("指定的地址无效....")
		os.Exit(0)
	}

}


func (cli *CLI)printChains(nodeID string){
	bc:=GetBlockchainObject(nodeID)
	if bc == nil{
		fmt.Println("没有区块可以打印。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.PrintChains()
}

//查询余额
func (cli *CLI)getBalance(address string,nodeID string){
	fmt.Println("查询余额：",address)
	bc := GetBlockchainObject(nodeID)

	if bc == nil{
		fmt.Println("数据库不存在，无法查询。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	//txOutputs:= bc.UnUTXOs(address)
	//for i,out:=range txOutputs{
	//	fmt.Println(i,"---->",out)
	//}
	//balance:=bc.GetBalance(address,[]*Transaction{})
	utxoSet:=&UTXOSet{bc}
	balance:= utxoSet.GetBalance(address)
	fmt.Printf("%s,一共有%d个Token\n",address,balance)
}


func (cli *CLI) createWallet(nodeID string){
	wallets:= NewWallets(nodeID)
	wallets.CreateNewWallet(nodeID)
}

func (cli *CLI) createGenesisBlockchain(address string,nodeID string){
	//fmt.Println(data)
	CreateBlockChainWithGenesisBlock(address,nodeID)


	bc:=GetBlockchainObject(nodeID)
	defer bc.DB.Close()
	if bc != nil{
		utxoSet:=&UTXOSet{bc}
		utxoSet.ResetUTXOSet()
	}

}

func (cli *CLI)addressLists(nodeID string){
	fmt.Println("打印所有的钱包地址。。")
	//获取
	Wallets:=NewWallets(nodeID)
	for address,_ := range Wallets.WalletsMap{
		fmt.Println("address:",address)
	}
}

