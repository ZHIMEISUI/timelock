package main

import (
	"github.com/timelock/lib"
	"os"
	"flag"
)

func main() {
	txFlag := "FundingTx"
	ChannelVersion := 0
	flag.StringVar(&txFlag, "t", "FundingTx", "set transactin type")
	flag.IntVar(&ChannelVersion, "cv", 0, "set channel version")
	flag.Parse()
	err := Execute(txFlag, ChannelVersion)
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
