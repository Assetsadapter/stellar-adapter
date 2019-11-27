package main

import (
	"github.com/blocktree/algorand-adapter/triam"
	"github.com/blocktree/openwallet/log"
)

var (
	WalletManager triam.WalletManager
)

func init() {
	wm := triam.WalletManager{}
	wm.Config = triam.NewConfig(triam.Symbol)
	wm.Blockscanner = triam.NewAlgoBlockScanner(&wm)
	wm.Decoder = triam.NewAddressDecoder(&wm)
	wm.TxDecoder = triam.NewTransactionDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	WalletManager = wm
}

func main() {

}
