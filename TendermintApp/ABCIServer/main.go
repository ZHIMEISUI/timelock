package main

import (
	"TimeLock/lib"
	"os"
)

func main() {
	err := Execute()
	if err != nil {
		lib.Log.Error(err)
		os.Exit(1)
	}
}
