package main

import (
	"github.com/timelock/lib"
	"os"
	"flag"
)

func main() {
	txFlag := "FundingTx"
	ChannelVersion := 0
	var Coins float64
	Coins = 0.0
	var From int64
	From = 0
	flag.StringVar(&txFlag, "t", txFlag, "set transactin type")
	flag.IntVar(&ChannelVersion, "cv", ChannelVersion, "set channel version")
	flag.Float64Var(&Coins, "coins", Coins, "set deposit coins")
	flag.Int64Var(&From, "from", From, "set previous transaction id")
	flag.Parse()
	err := Execute(txFlag, From, uint8(ChannelVersion), Coins)
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
