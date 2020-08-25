package timelock

import (

	"os"
	"io"
	"fmt"
	"strconv"
	// "bytes"
	// "io/ioutil"
	"strings"
	"encoding/json"
	"encoding/binary"


	"github.com/timelock/lib"
	"github.com/timelock/controllers"

	dbm "github.com/tendermint/tm-db"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/abci/example/code"
	// cmn "github.com/tendermint/tendermint/tmlibs/common"
)

var (
	stateKey        = []byte("stateKey")
	txRecord		= controllers.Transaction{}

	ProtocolVersion uint64 = 0x1
)

type State struct{
	DB 		dbm.DB						`bson:"db"			json:"db"`
	Height 	int64						`bson:"height"		json:"height"`
	AppHash	[]byte						`bson:"app_hash"	json:"app_hash"`
	Size 	int64  						`bson:"size"		json:"size"`
	Tx 		controllers.Transaction		`bson:"tx"			json:"tx"`
}

func loadState(db dbm.DB) State{
	var state State
	state.DB = db
	stateBytes, err := db.Get(stateKey)
	lib.Log.Debug("stateKey: ")
	lib.Log.Debug(string(stateKey))
	if err != nil {
		lib.Log.Error(err)
		panic(err)
	}
	if len(stateBytes) == 0 {
		lib.Log.Error("stateBytes is nil...")
		return state
	}
	err = json.Unmarshal(stateBytes, &state)
	if err != nil {
		panic(err)
	}
	return state
}

func setStateTx(txmap map[string]string, app *TimelockApplication){
	app.state.Tx.ID, _ = strconv.ParseInt(txmap["ID"], 10, 64)
	app.state.Tx.Flag = txmap["Flag"]
	// app.state.Tx.Height, _ = strconv.ParseUint(txmap["Height"], 10, 64)
	app.state.Tx.PreTxId, _ = strconv.ParseInt(txmap["PreTxId"], 10, 64)
	app.state.Tx.From = txmap["From"]
	app.state.Tx.To = txmap["To"]
	coin,_ := strconv.ParseFloat(txmap["Coin"], 32)
	app.state.Tx.Coin = float32(coin)
	ncommit,_ := strconv.ParseUint(txmap["NCommit"], 10, 8)
	app.state.Tx.NCommit = uint8(ncommit)
	app.state.Tx.Sig = txmap["Sig"]
	lib.Log.Notice(app.state.Tx)
}
func clearTx(app *TimelockApplication)  {
	app.state.Tx.ID, _ = strconv.ParseInt("", 10, 64)
	app.state.Tx.Flag = ""
	// app.state.Tx.Height, _ = strconv.ParseUint("", 10, 64)
	app.state.Tx.From = ""
	app.state.Tx.To = ""
	coin,_ := strconv.ParseFloat("", 32)
	app.state.Tx.Coin = float32(coin)
	ncommit,_ := strconv.ParseUint("", 10, 8)
	app.state.Tx.NCommit = uint8(ncommit)
	app.state.Tx.Sig = ""
}

func saveState(app *TimelockApplication) {
	stateBytes, err := json.Marshal(app.state)
	if err != nil {
		panic(err)
	}
	err = app.state.DB.Set(stateKey, stateBytes)
	if err != nil {
		lib.Log.Error("state.DB.Set() err: ")
		lib.Log.Error(err)
		panic(err)
	}
}

func txHandle(tx string) map[string]string {
	txhandle := strings.Replace(tx, "'", "", -1)
	txhandle = strings.Replace(string(txhandle), "{", "", -1)
	txhandle = strings.Replace(string(txhandle), "[", "", -1)
	txhandle = strings.Replace(string(txhandle), "]", "", -1)
	txhandle = strings.Replace(string(txhandle), "}", "", -1)
	txs := strings.Split(string(txhandle), ",")
	txmap := make(map[string]string)
	
	for _ , t := range txs {
		tsplit := strings.Split(string(t), ":")
		txmap[tsplit[0]] = tsplit[1]
	}
	return txmap
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

func TriggerTxVerify(app *TimelockApplication, tx map[string]string, f *os.File) bool {
	// b,_ := app.state.DB.Has([]byte(strconv.FormatInt(app.state.Tx.PreTxId, 10)))
	// lib.Log.Notice(b)
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

func NewTimelockApplication(memDB dbm.DB) *TimelockApplication {
	lib.Log.Debug("NewTimelockApplication")
	state := loadState(memDB)
	return &TimelockApplication{state: state}
}

func (app *TimelockApplication) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	lib.Log.Debug("Info")
	return types.ResponseInfo{Data: fmt.Sprintf("Info(): TimeLock Test")}
}

func (app *TimelockApplication) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	lib.Log.Debug("DeliverTx")
	lib.Log.Notice(string(req.Tx))

	txmap:= txHandle(string(req.Tx))

	lib.Log.Debug("app.state: ")
	lib.Log.Debug(app.state)
	if txmap["Flag"] == "FundingTx" {
		logTx("DeliverTx", txmap)
		statejson, _ := json.Marshal(app.state)
		lib.Log.Debug(string(statejson))
		if !FundingTxVerify(txmap) {
			lib.Log.Warning("Code: "+strconv.FormatUint(uint64(code.CodeTypeBadNonce), 10))
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}

		f, err := os.OpenFile("./log/timelock.db/timelock.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil{
			lib.Log.Warning("write timelock.txt error!")
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		txline, err := f.Write([]byte(string(req.Tx)+"\n"))
		lib.Log.Notice(txline)
		defer f.Close()

	} else if txmap["Flag"] == "TriggerTx" {
		logTx("DeliverTx", txmap)
		f, err := os.OpenFile("./log/timelock.db/timelock.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil{
			lib.Log.Warning("write timelock.txt error!")
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}

		if !TriggerTxVerify(app, txmap, f) {
			lib.Log.Warning("Code: "+strconv.FormatUint(uint64(code.CodeTypeBadNonce), 10))
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		
		txline, err := f.Write([]byte(string(req.Tx)+"\n"))
		lib.Log.Notice(txline)
		defer f.Close()
	} else if txmap["Flag"] == "SettlementTx" {
		logTx("DeliverTx", txmap)
		if !SettlementTxVerify(txmap) {
			lib.Log.Warning("Code: "+strconv.FormatUint(uint64(code.CodeTypeBadNonce), 10))
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		f, err := os.OpenFile("./log/timelock.db/timelock.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil{
			lib.Log.Warning("write timelock.txt error!")
			return types.ResponseDeliverTx{Code: code.CodeTypeBadNonce}
		}
		txline, err := f.Write([]byte(string(req.Tx)+"\n"))
		lib.Log.Notice(txline)
		defer f.Close()
	}
	setStateTx(txmap, app)
	lib.Log.Debug("app.state.Tx ---> ")
	lib.Log.Debug(app.state.Tx)

	events :=  []types.Event{
		{
			Type: "app",
			Attributes: []types.EventAttribute{
				{Key:[]byte("Transaction Type"), Value:[]byte(txmap["Flag"]), Index:true},
				{Key:[]byte("Previous Transaction ID"), Value:[]byte(txmap["PreTxId"]), Index:true},
				{Key:[]byte("Transaction ID"), Value:[]byte(txmap["ID"]), Index:true},
				{Key:[]byte("Blockheight"), Value:[]byte(strconv.FormatInt(app.state.Height, 10)), Index:true},
				{Key:[]byte("From"), Value:[]byte(txmap["From"]), Index:true},
				{Key:[]byte("To"), Value:[]byte(txmap["To"]), Index:true},
				{Key:[]byte("Deposit Coins"), Value:[]byte(txmap["Coin"]), Index:true},
				{Key:[]byte("Channel Version"), Value:[]byte(txmap["NCommit"]), Index:true},
				{Key:[]byte("Sig"), Value:[]byte(txmap["Sig"]), Index:true},
			},
		},
	}
	return types.ResponseDeliverTx{Code: code.CodeTypeOK, Events: events}
}


func (app *TimelockApplication) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	lib.Log.Debug("CheckTx")
	lib.Log.Notice(string(req.Tx))
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}

func (app *TimelockApplication) Commit() types.ResponseCommit {
	lib.Log.Debug("Commit")
	appHash := make([]byte, 8)
	binary.PutVarint(appHash, app.state.Size)
	app.state.AppHash = appHash
	app.state.Height++
	saveState(app)

	stateDBjson, _ := json.Marshal(app.state.DB)
	lib.Log.Debug("stateDBjson: "+string(stateDBjson))
	statejson, errs := json.Marshal(app.state)
	lib.Log.Debug("statejson: "+string(statejson))
	clearTx(app)
	if errs != nil {return types.ResponseCommit{}}
	return types.ResponseCommit{Data: statejson}
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