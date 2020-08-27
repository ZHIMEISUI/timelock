package main

import (
	"github.com/timelock/controllers"
	"github.com/timelock/lib"
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	// "math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func Execute(flag string, From int64, ChannelVersion uint8, Coins float64, SecretT int64) error {
	lib.Log.Notice("Starting Bob UI Client... ")

	f, err := os.Create("logs/client.log")
	if err != nil {
		fmt.Println("Client log init error:", err)
	}
	multiWriter := io.MultiWriter(f, os.Stdout)

	go func() {
		cmd := exec.Command("bash", "-c", "sh run-client.sh")
		cmd.Stdout = multiWriter
		cmd.Start()
	}()


	go func() {
		blocksNumber := 1                                     // how many blocks
		transactionsPerBlock := 1                            // how many transactions in each block
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		
		for i := 0; i < blocksNumber; i++ {
			time.Sleep(time.Second * 1)
			transactions := []controllers.Transaction{}
			tran := controllers.Transaction{}

			for j := 0; j < transactionsPerBlock; j++ {
				if flag == "FundingTx"{
					_, _ = tran.CreateFundingTx(From, "Bob", float32(Coins), "BobSig")
				}else if flag == "TriggerTx"{
					_, _ = tran.CreateTriggerTx(From, "Alice&&Bob", float32(Coins), ChannelVersion, "BobSig")
				}else if flag == "SettlementTx"{
					_, _ = tran.CreateSettlementTx(From, "Bob", float32(Coins), ChannelVersion, SecretT, "BobSig")
				}
				transactions = append(transactions, tran)
			}
			bytes, _ := json.Marshal(&transactions)
			data := strings.Replace(string(bytes), "\"", "'", -1)

			tx := data

			tmCommit(tx)
		}
	}()

	runConsole()

	return nil
}

func tmAsync(tx string) {
	// url := "http://localhost:46657/broadcast_tx_async?tx=\"" + tx + "\""
	url := "http://localhost:26658/broadcast_tx_async?tx=\"" + tx + "\""
	txHandle(url)
}

func tmSync(tx string) {
	// url := "http://localhost:46657/broadcast_tx_sync?tx=\"" + tx + "\""
	url := "http://localhost:26658/broadcast_tx_sync?tx=\"" + tx + "\""
	txHandle(url)
}

func tmCommit(tx string) {
	// fmt.Printf("szm prints tx in tmCommit()...: %s \n", tx)
	// url := "http://localhost:46657/broadcast_tx_async?tx=\"" + tx + "\""
	url := "http://localhost:26657/broadcast_tx_async?tx=\"" + tx + "\""
	txHandle(url)
}

func txHandle(url string) {
	lib.Log.Debug("szm debugs url in txHandle()...: "+url)
	fmt.Printf("\n")
	resp, err := http.Get(url)
	fmt.Printf("szm log resp in txHandle()...: ")
	fmt.Println(resp)
	if err != nil{
		lib.HandleError(err)
	}
	fmt.Printf("\n")

	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("szm prints in txHandle()...:")
		lib.HandleError(err)
		fmt.Printf("\n")

		var data interface{}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(body, &data)
		lib.Log.Notice(data)
	}
	
}

func runConsole() error {
	for {
		fmt.Printf("> ...")
		bufReader := bufio.NewReader(os.Stdin)
		line, more, err := bufReader.ReadLine()
		if more {
			return errors.New("Input is too long")
		} else if err != nil {
			return err
		}

		fmt.Println("Client,", line)
	}
}
