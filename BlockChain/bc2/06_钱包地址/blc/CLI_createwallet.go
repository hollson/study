package blc

func (cli *CLI) createWallet() {
	wallets := NewWallets()
	wallets.CreateNewWallet()
}
