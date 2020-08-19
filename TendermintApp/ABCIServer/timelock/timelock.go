package timelock

import (

	"fmt"
	"strconv"
	// "bytes"
	"strings"

	"github.com/timelock/lib"

	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/abci/example/code"
	// cmn "github.com/tendermint/tendermint/tmlibs/common"
)

var _ types.Application = (*TimelockApplication)(nil)

type TimelockApplication struct {
	types.BaseApplication
	tx_type string
	flag bool
}

func NewTimelockApplication() *TimelockApplication {
	lib.Log.Debug("NewTimelockApplication")
	flag := true
	return &TimelockApplication{flag: flag}
}

func (app *TimelockApplication) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	lib.Log.Debug("Info")
	return types.ResponseInfo{Data: fmt.Sprintf("TimeLock Test")}
}

func (app *TimelockApplication) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	lib.Log.Debug("DeliverTx")
	lib.Log.Notice(string(req.Tx))
	// lib.Log.Notice(req.Tx)
	txhandle := strings.Replace(string(req.Tx), "'", "", -1)
	txhandle = strings.Replace(string(txhandle), "{", "", -1)
	txhandle = strings.Replace(string(txhandle), "[", "", -1)
	lib.Log.Debug(txhandle)
	txs := strings.Split(string(txhandle), ",")
	txmap := make(map[string]string)
	for _ , t := range txs {
		tsplit := strings.Split(string(t), ":")
		txmap[tsplit[0]] = tsplit[1]
	}
	if txmap["Flag"] == "FundingTx" {
		lib.Log.Debug("From: "+txmap["From"])
		lib.Log.Debug("To: "+txmap["To"])
		return types.ResponseDeliverTx{Code: code.CodeTypeOK}
	} else if txmap["Flag"] == "TriggerTx" {
		return types.ResponseDeliverTx{Code: code.CodeTypeOK}
	} else if txmap["Flag"] == "SettlementTx" {
		return types.ResponseDeliverTx{Code: code.CodeTypeOK}
	}
	
	lib.Log.Debug()
	return types.ResponseDeliverTx{Code: code.CodeTypeUnknownError}
}


func (app *TimelockApplication) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	lib.Log.Debug("CheckTx")
	lib.Log.Notice(string(req.Tx))
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}

func (app *TimelockApplication) Commit() types.ResponseCommit {
	lib.Log.Debug("Commit")
	// Save a new version
	var flag bool
	flag = app.flag

	if app.flag{
		lib.Log.Notice("flag",flag)
	}

	lib.Log.Debug("timelock flag", flag)
	return types.ResponseCommit{Data: []byte(strconv.FormatBool(flag))}
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