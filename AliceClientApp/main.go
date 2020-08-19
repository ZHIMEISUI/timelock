package main

import (
	"github.com/timelock/lib"
	"os"
	"flag"
)

func main() {
	txFlag := "FundingTx"
	ChannelVersion := 0
	var Coins float32
	Coins = 0.0
	flag.StringVar(&txFlag, "t", "FundingTx", "set transactin type")
	flag.IntVar(&ChannelVersion, "cv", 0, "set channel version")
	flag.Var(&Coins, "coins", 0, "set deposit coins")
	flag.Parse()
	err := Execute(txFlag, uint8(ChannelVersion), Coins)
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
