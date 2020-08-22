package timelock

import (

	"fmt"
	"strconv"
	// "bytes"
	"strings"
	"encoding/json"


	"github.com/timelock/lib"
	"github.com/timelock/controllers"

	dbm "github.com/tendermint/tm-db"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/abci/example/code"
	// cmn "github.com/tendermint/tendermint/tmlibs/common"
)

var (
	stateKey        = []byte("stateKey")

	ProtocolVersion uint64 = 0x1
)

type State struct{
	DB dbm.DB
	Height int		`bson:"height"	json:"height"`
	AppHash	[]byte	`json:"app_hash"`
	Tx controllers.Transaction
}

func loadState(db dbm.DB) State{
	var state State
	state.DB = db
	stateBytes, err := db.Get(stateKey)
	if err != nil {
		panic(err)
	}
	if len(stateBytes) == 0 {
		return state
	}
	err = json.Unmarshal(stateBytes, &state)
	if err != nil {
		panic(err)
	}
	return state
}

func saveState(state State) {
	stateBytes, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	err = state.DB.Set(stateKey, stateBytes)
	if err != nil {
		panic(err)
	}
}


func logTx(funcname string, txmap map[string]string){
	lib.Log.Debug(funcname+" starts Debug...")
	lib.Log.Debug("Transaction ID: "+txmap["ID"])
	lib.Log.Debug("Transaction Type: "+txmap["Flag"])
	lib.Log.Debug("Current Time: "+txmap["CurrentTime"])
	lib.Log.Debug("From: "+txmap["From"])
	lib.Log.Debug("To: "+txmap["To"])
	lib.Log.Debug("Deposit Coins: "+txmap["Coin"])
	lib.Log.Debug("Channel Version: "+txmap["NCommit"])
	lib.Log.Debug("Sig: "+txmap["Sig"])
	lib.Log.Debug(funcname+" ends Debug...")
}

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

func TriggerTxVerify(tx map[string]string) bool {
	if tx["Flag"] == "TriggerTx"{
		from,_ := tx["From"]
		if from != "Alice&&Bob" {
			lib.Log.Warning("Your Trigger Transaction is not valid")
			lib.Log.Warning(tx["Flag"]+" should send to both Alice and Bob!")
			return false
		}
		lib.Log.Notice("Your Trigger Transaction is recorded successfully!")
		return true
	}
	return false
}

func SettlementTxVerify(tx map[string]string) bool {
	if tx["Flag"] == "SettlementTx"{
		from,_ := tx["From"]
		if from != "Alice&&Bob" {
			lib.Log.Warning("Your Settlement Transaction is not valid")
			lib.Log.Warning(tx["Flag"]+" should send from both Alice and Bob!")
			return false
		}
		lib.Log.Notice("Your Settlement Transaction is recorded successfully!")
		return true
	}
	return false
}


// -------------------------------------------------------------------
var _ types.Application = (*TimelockApplication)(nil)

type TimelockApplication struct {
	types.BaseApplication
	
	state State
}

func NewTimelockApplication() *TimelockApplication {
	lib.Log.Debug("NewTimelockApplication")
	state := loadState(dbm.NewMemDB())
	return &TimelockApplication{state: state}
}

func (app *TimelockApplication) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	lib.Log.Debug("Info")
	return types.ResponseInfo{Data: fmt.Sprintf("Info(): TimeLock Test")}
}

func (app *TimelockApplication) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	lib.Log.Debug("DeliverTx")
	lib.Log.Notice(string(req.Tx))

	txhandle := strings.Replace(string(req.Tx), "'", "", -1)
	txhandle = strings.Replace(string(txhandle), "{", "", -1)
	txhandle = strings.Replace(string(txhandle), "[", "", -1)
	txhandle = strings.Replace(string(txhandle), "]", "", -1)
	txhandle = strings.Replace(string(txhandle), "}", "", -1)
	lib.Log.Debug(txhandle)
	txs := strings.Split(string(txhandle), ",")
	txmap := make(map[string]string)
	for _ , t := range txs {
		tsplit := strings.Split(string(t), ":")
		txmap[tsplit[0]] = tsplit[1]
	}
	if txmap["Flag"] == "FundingTx" {
		logTx("DeliverTx", txmap)
		statejson, _ := json.Marshal(app.state)
		lib.Log.Debug(string(statejson))
		if !FundingTxVerify(txmap) {
			lib.Log.Warning("Code: "+strconv.FormatUint(uint64(code.CodeTypeBadNonce), 10))
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		return types.ResponseDeliverTx{Code: code.CodeTypeOK}
	} else if txmap["Flag"] == "TriggerTx" {
		logTx("DeliverTx", txmap)
		if !TriggerTxVerify(txmap) {
			lib.Log.Warning("Code: "+strconv.FormatUint(uint64(code.CodeTypeBadNonce), 10))
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		return types.ResponseDeliverTx{Code: code.CodeTypeOK}
	} else if txmap["Flag"] == "SettlementTx" {
		logTx("DeliverTx", txmap)
		if !SettlementTxVerify(txmap) {
			lib.Log.Warning("Code: "+strconv.FormatUint(uint64(code.CodeTypeBadNonce), 10))
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		return types.ResponseDeliverTx{Code: code.CodeTypeOK}
	}
	
	return types.ResponseDeliverTx{Code: code.CodeTypeUnknownError}
}


func (app *TimelockApplication) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	lib.Log.Debug("CheckTx")
	lib.Log.Notice(string(req.Tx))
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}

func (app *TimelockApplication) Commit() types.ResponseCommit {
	lib.Log.Debug("Commit")
	app.state.Height++
	saveState(app.state)
	statejson, _ := json.Marshal(app.state)
	lib.Log.Debug(string(statejson))
	txjson, errs := json.Marshal(app.state.Tx)
	lib.Log.Debug(string(txjson))
	if errs != nil {return types.ResponseCommit{}}
	return types.ResponseCommit{Data: txjson}
}

func (app *TimelockApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
	lib.Log.Debug("Query")
	// switch resQuery.Path {
	// case "flag":
	// 	return types.ResponseQuery{Value: []byte(cmn.Fmt("%t", app.flag))}
	// default:
	// 	return types.ResponseQuery{Log: cmn.Fmt("Invalid query path. Expected hash or tx, got %v", reqQuery.Path)}
	// }
	return 
}