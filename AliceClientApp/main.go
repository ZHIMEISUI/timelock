package main

import (
	"github.com/timelock/lib"
	"os"
	"flag"
)

func main() {
	var txFlag string
	flag.StringVar(&txFlag, "t", "FundingTx", "set transactin type")
	flag.Parse()
	err := Execute(txFlag)
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
