package timelock

import (

	"fmt"

	"github.com/timelock/lib"

	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/abci/example/code"
	cmn "github.com/tendermint/tendermint/tmlibs/common"
)

var _ types.Application = (*TimelockApplication)(nil)

type TimelockApplication struct {
	types.BaseApplication

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
RequestDeliverTx
func (app *TimelockApplication) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	lib.Log.Debug("DeliverTx")
	lib.Log.Notice(string(req.Tx))
	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}

// func (app *TimelockApplication) CheckTx(tx []byte) types.ResponseCheckTx {
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
	return types.ResponseCommit{Data: "true"}
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