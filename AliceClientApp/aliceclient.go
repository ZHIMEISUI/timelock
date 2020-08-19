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

func Execute(flag string) error {
	lib.Log.Notice("Starting Alice UI Client... ")

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
		blocksNumber := 5                                     // how many blocks
		transactionsPerBlock := 10                            // how many transactions in each block
		// players := []string{"Lei", "Jack", "Pony", "Richard"} // 4 players
		// random := rand.New(rand.NewSource(time.Now().UnixNano()))
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		
		for i := 0; i < blocksNumber; i++ {
			time.Sleep(time.Second * 1)
			transactions := []controllers.Transaction{}
			tran := controllers.Transaction{}

			for j := 0; j < transactionsPerBlock; j++ {
				if flag == "FundingTx"{
					_, _ = tran.CreateFundingTx("Alice", "Alice&&Bob", 5, "ChannelVersion", "AliceSig")
				}else if flag == "TriggerTx"{
					_, _ = tran.CreateTriggerTx("Alice&&Bob", "Alice&&Bob", 5, "ChannelVersion", "AliceSig")
				}else if flag == "SettlementTx"{
					_, _ = tran.CreateSettlementTx("Alice&&Bob", "Alice", 5, "ChannelVersion", "AliceSig")
				}
				transactions = append(transactions, tran)
			}
			fmt.Printf("szm log transactions in go func()...: %s \n", transactions)
			fmt.Printf("szm log transactions type in go func()...: %T \n", transactions)

			bytes, _ := json.Marshal(&transactions)
			// fmt.Printf("szm log bytes in go func()...: ")
			// fmt.Println(bytes)
			data := strings.Replace(string(bytes), "\"", "'", -1)
			// lib.Log.Notice("szm log data in go func()...:"+data)
			// fmt.Printf("szm log data type in go func()...: %T \n", data)
			// fmt.Printf("\n")

			// tx := "id=" + lib.Int64ToString(tran.ID) + "&flag=" + tran.flag
			tx := data
			fmt.Printf("szm log tx in go func()...: %s \n", tx)
			fmt.Printf("szm log tx type in go func()...: %T \n", tx)
			
			// tmAsync(tx)
			// tmCommit(lib.Int64ToString(transactions[0].ID))
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
	fmt.Printf("szm prints tx in tmCommit()...: %s \n", tx)
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
