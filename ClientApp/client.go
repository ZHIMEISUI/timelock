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

func Execute() error {
	lib.Log.Notice("Starting UI Client")

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

			for j := 0; j < transactionsPerBlock; j++ {
				// from := players[random.Intn(len(players))]
				// to := players[random.Intn(len(players))]
				// for from == to {
				// 	to = players[random.Intn(len(players))]
				// }
				// btc := float32(random.Intn(10) + 1)
				flag := true

				// tran := controllers.Transaction{
				// 	From:    from,
				// 	To:      to,
				// 	Bitcoin: btc,
				// }
				tran := controllers.Transaction{
					Flag:    flag,
				}
				_, _ = tran.Create()
				transactions = append(transactions, tran)
			}
			fmt.Printf("szm log...%s", transactions)
			bytes, _ := json.Marshal(&transactions)
			data := strings.Replace(string(bytes), "\"", "'", -1)
			lib.Log.Notice("szm log...:"+data)

			tx := "id=" + lib.Int64ToString(tran.ID) + "&flag=" + tran.flag
			// tx := data
			fmt.Printf("szm log...%s", tx)
			// tmAsync(tx)
			tmCommit(tx)
		}
	}()

	runConsole()

	return nil
}

func tmAsync(tx string) {
	url := "http://localhost:46657/broadcast_tx_async?tx=\"" + tx + "\""
	txHandle(url)
}

func tmSync(tx string) {
	url := "http://localhost:46657/broadcast_tx_sync?tx=\"" + tx + "\""
	txHandle(url)
}

func tmCommit(tx string) {
	url := "http://localhost:46657/broadcast_tx_async?tx=\"" + tx + "\""
	txHandle(url)
}

func txHandle(url string) {
	lib.Log.Debug("szm log..."+url)
	resp, err := http.Get(url)
	lib.HandleError("szm log err..."+err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	lib.HandleError("szm log err..."+err)

	var data interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &data)
	lib.Log.Notice(data)
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
