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
	Coins = 10.0
	var From, Secrett int64
	From, Secrett = 0, 0
	// flag.StringVar(&txFlag, "t", txFlag, "set transactin type")
	// flag.IntVar(&ChannelVersion, "cv", ChannelVersion, "set channel version")
	// flag.Float64Var(&Coins, "coins", Coins, "set deposit coins")
	// flag.Int64Var(&From, "from", From, "set previous transaction id")
	// flag.Int64Var(&Secrett, "st", Secrett, "set the owner's secret T")
	flag.Parse()
	err := Execute(txFlag, From, uint8(ChannelVersion), Coins, Secrett)
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
