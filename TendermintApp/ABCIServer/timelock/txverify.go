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

		// lib.Log.Notice(string(chunk))
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

		// lib.Log.Notice(string(chunk))
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
		// previous/trigger tx info
		txmap := lib.TxHandle(txstring)
		tgbh, _ := strconv.ParseUint(txmap["BlockHeight"],10,8)
		tgtl, _ := strconv.ParseUint(txmap["TimeLock"],10,8)
		tgnc, _ := strconv.ParseUint(txmap["NCommit"],10,8)

		//this settlement tx info
		stbh, _ := strconv.ParseUint(tx["BlockHeight"],10,8)
		stnc, _ := strconv.ParseUint(tx["NCommit"],10,8)

		if uint8(stbh) <= uint8(tgbh)+uint8(tgtl) {
			if uint8(stnc) > uint8(tgnc) { // 若另一方提供更高版本的NCommit
				// 该交易owner(不同于TriggerTx的owner)可以拿走全部deposit
				if tx["Sig"] != txmap["Sig"]{
					lib.Log.Notice("Settlement Scenario 1")
					lib.Log.Notice(tx["Sig"]+" provides a higher version and takes all coins.")
					lib.Log.Notice("Your Settlement Transaction is recorded successfully!")
					
					return true
				}
				lib.Log.Warning(tx["Sig"])
				lib.Log.Warning(txmap["Sig"])
				lib.Log.Warning("Settlement 1 failed")
			} else { // 若另一方不提供更高版本的NCommit
				// 验证t_alice
				// 该交易owner(不同于triggerTx的owner)分配FundingTx中的钱给双方
				if tx["Sig"] != txmap["Sig"]{
					lib.Log.Notice("Settlement Scenario 2")
					lib.Log.Notice("Your Settlement Transaction is recorded successfully! Each party retrieves their funding coins.")
					
					return true
				}
				lib.Log.Warning(tx["Sig"])
				lib.Log.Warning(txmap["Sig"])
				lib.Log.Warning("Settlement 2 failed")
			}
		} else {
			// 该交易owner(与TriggerTx的owner一致)可以拿走全部deposit
			if tx["Sig"] == txmap["Sig"]{
				lib.Log.Notice("Settlement Scenario 3")
				lib.Log.Notice(tx["Sig"]+" fails to release his/her secret. So "+ txmap["Sig"]+" takes all coins.")
				lib.Log.Notice("Your Settlement Transaction is recorded successfully!")
				
				return true
			}
			lib.Log.Warning(tx["Sig"])
			lib.Log.Warning(txmap["Sig"])
			lib.Log.Warning("Settlement 3 failed")
		}
	}
	lib.Log.Warning("Your Settlement Transaction is not valid")
	return false
}