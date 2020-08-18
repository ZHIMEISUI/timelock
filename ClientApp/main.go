package main

import (
	"github.com/timelock/lib"
	"os"
)

func main() {
	err := Execute()
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
	err = ExecuteAlice()
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
	err = ExecuteBob()
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
