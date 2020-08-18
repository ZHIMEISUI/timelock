package controllers

import "github.com/timelock/lib"

/*
Transaction
*/
type Transaction struct {
	ID      int64   `bson:"id" json:"id"`
	Flag string `bson:"flag" json:"flag"`
	CurrentTime uint8 `bson:"currenttime" json:"currenttime"`
	From    string  `bson:"from" json:"from"`
	To      string  `bson:"to" json:"to"`
	Coin float32 `bson:"coin" json:"coin"`
	NCommit string `bson:"ncommit" json:"ncommit"`
	Sig string `bson:"ncommit" json:"ncommit"`
}

type FundingTransaction struct {
	Transaction
}

type TriggerTransaction struct {
	Transaction
}

type SettlementTransaction struct {
	Transaction
}

/*
Create :creating transactions
*/
func (t *Transaction) CreateTx() (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.Flag = "Transaction"
	lib.Log.Debug("Create Transaction:", t)
	return true, nil
}

func (t *FundingTransaction) CreateFundingTx() (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.flag = "FundingTx"
	lib.Log.Debug("Create Funding Transaction:", t)
	return true, nil
}

func (t *TriggerTransaction) CreateTriggerTx() (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.flag = "TriggerTx"
	lib.Log.Debug("Create Trigger Transaction:", t)
	return true, nil
}

func (t *SettlementTransaction) CreateSettlementTx() (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.flag = "SettlementTx"
	lib.Log.Debug("Create Settlement Transaction:", t)
	return true, nil
}
