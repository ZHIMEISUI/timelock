package timelock

import (

	"os"
	"io"
	"fmt"
	"strconv"
	// "bytes"
	// "io/ioutil"
	"strings"
	// "encoding/json"
	// "encoding/binary"


	"github.com/timelock/lib"
	// "github.com/timelock/controllers"

	// dbm "github.com/tendermint/tm-db"
	// "github.com/tendermint/tendermint/abci/types"
	// "github.com/tendermint/tendermint/abci/example/code"
	// cmn "github.com/tendermint/tendermint/tmlibs/common"
)

func SigVerify() bool {return true}

func NCommitVerify() bool {return true}

func FundingTxVerify(tx map[string]string) bool {
	if tx["Flag"] == "FundingTx"{
		coin,_ := strconv.Atoi(tx["Coin"])
		if coin <= 0 {
			lib.Log.Warning("Your Funding Transaction is not valid")
			lib.Log.Warning(tx["From"]+" deposits coin is: "+ tx["Coin"] + ". The expected deposits in Funding Transaction is higher than 0.")
			return false
		}
		lib.Log.Notice("Your Funding Transaction is recorded successfully!")
		return true
	}
	return false
}

func TriggerTxVerify(app *TimelockApplication, tx map[string]string, f *os.File) bool {
	if tx["Flag"] == "TriggerTx"{
		var chunk []byte
		buf := make([]byte, 1024)

		for {
			//从file读取到buf中
			n, err := f.Read(buf)
			if err != nil && err != io.EOF{
				fmt.Println("read buf fail", err)
				return false
			}
			//说明读取结束
			if n == 0 {
				break
			}
			//读取到最终的缓冲区中
			chunk = append(chunk, buf[:n]...)
		}

		lib.Log.Notice(string(chunk))
		txs := strings.Split(string(chunk), "***")
		

		txstring, b := lib.Has(txs, tx["From"], "ID")
		if !b {
			lib.Log.Warning("Your Trigger Transaction is not valid")
			return false
		}
		var txarray []string
		txarray = append(txarray, txstring)
		if _, b = lib.Has(txarray, "FundingTx", "Flag"); !b {
			lib.Log.Warning("Your Trigger Transaction is not valid")
			return false
		}

			lib.Log.Notice("Your Trigger Transaction is recorded successfully!")
			return true
	}
	return false
}

func SettlementTxVerify(app *TimelockApplication, tx map[string]string, f *os.File) bool {
	if tx["Flag"] == "SettlementTx"{

		var chunk []byte
		buf := make([]byte, 1024)

		for {
			//从file读取到buf中
			n, err := f.Read(buf)
			if err != nil && err != io.EOF{
				fmt.Println("read buf fail", err)
				return false
			}
			//说明读取结束
			if n == 0 {
				break
			}
			//读取到最终的缓冲区中
			chunk = append(chunk, buf[:n]...)
		}

		lib.Log.Notice(string(chunk))
		txs := strings.Split(string(chunk), "***")
		// from := strconv.FormatInt(app.state.Tx.From, 10)
		

		txstring, b := lib.Has(txs, tx["From"], "ID")
		if !b {
			lib.Log.Warning("Your Settlement Transaction is not valid")
			return false
		}
		var txarray []string
		txarray = append(txarray, txstring)
		if txstring, b = lib.Has(txarray, "TriggerTx", "Flag"); !b {
			lib.Log.Warning("Your Settlement Transaction is not valid")
			return false
		}
		txmap := lib.TxHandle(txstring)
		bh, _ := strconv.ParseUint(txmap["BlockHeight"],10,8)
		tl, _ := strconv.ParseUint(txmap["TimeLock"],10,8)
		nc, _ := strconv.ParseUint(txmap["NCommit"],10,8)
		if app.state.Height <= uint8(bh)+uint8(tl) {
			if app.state.Tx.NCommit > uint8(nc) { // 若另一方提供更高版本的NCommit
				// 该交易owner(不同于TriggerTx的owner)可以拿走全部deposit
				if app.state.Tx.Sig != txmap["Sig"]{
					lib.Log.Notice("Your Settlement Transaction is recorded successfully!")
					lib.Log.Notice("Settlement 1")
					return true
				}
				lib.Log.Warning("Settlement 1 failed")
			} else { // 若另一方不提供更高版本的NCommit
				// 验证t_alice
				// 该交易owner(不同于triggerTx的owner)分配FundingTx中的钱给双方
				if app.state.Tx.Sig != txmap["Sig"]{
					lib.Log.Notice("Your Settlement Transaction is recorded successfully!")
					lib.Log.Notice("Settlement 2")
					return true
				}
				lib.Log.Warning("Settlement 2 failed")
			}
		} else {
			// 该交易owner(与TriggerTx的owner一致)可以拿走全部deposit
			if app.state.Tx.Sig == txmap["Sig"]{
				lib.Log.Notice("Your Settlement Transaction is recorded successfully!")
				lib.Log.Notice("Settlement 3")
				return true
			}
			lib.Log.Warning("Settlement 3 failed")
		}
	}
	lib.Log.Warning("Your Settlement Transaction is not valid")
	return false
}