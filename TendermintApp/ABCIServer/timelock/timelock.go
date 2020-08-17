package timelock

import (

	"TimeLock/lib"

	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/abci/example/code"
)

var _ types.Application = (*TimelockApplication)(nil)

type TimelockApplication struct {
	types.BaseApplication

	flag bool
}

func NewTimelockApplication() *TimelockApplication {
	lib.Log.Debug("NewTimelockApplication")
	flag := true
	return &TimelockApplication{state: flag}
}

func (app *TimelockApplication) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	lib.Log.Debug("Info")
	return types.ResponseInfo{Data: fmt.Sprintf("TimeLock Test"}
}

func (app *TimelockApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {
	lib.Log.Debug("DeliverTx")
	lib.Log.Notice(string(tx))
	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}

func (app *TimelockApplication) CheckTx(tx []byte) types.ResponseCheckTx {
	lib.Log.Debug("CheckTx")
	lib.Log.Notice(string(tx))
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}

func (app *TimelockApplication) Commit() types.ResponseCommit {
	lib.Log.Debug("Commit")
	// Save a new version
	var hash []byte
	var err error

	if app.flag{
		lib.Log.Notice("flag",flag)
	}

	lib.Log.Debug("timelock flag", flag)
	return types.ResponseCommit{Code: code.CodeTypeOK, Data: flag}
}

func (app *TimelockApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
	lib.Log.Debug("Query")
	resQuery.Log = "exists"
}