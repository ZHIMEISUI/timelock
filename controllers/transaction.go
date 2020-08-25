package controllers

import "github.com/timelock/lib"

/*
Transaction
*/
type Transaction struct {
	ID      	int64   `bson:"id" 				json:"id"`
	Flag 		string 	`bson:"flag" 			json:"flag"`
	// PreTxId		int64	`bson:"pretxid" 		json:"pretxid"`
	// Height  	uint64	`bson:"height" 			json:"height"`
	TimeLock	uint8	`bson:"timelock"		json:"timelock"`
	From    	int64  	`bson:"from" 			json:"from"`
	// From    	string  `bson:"from" 			json:"from"`
	To      	string  `bson:"to" 				json:"to"`
	Coin 		float32 `bson:"coin" 			json:"coin"`
	NCommit 	uint8 	`bson:"channelversion" 	json:"channelversion"`
	Sig 		string 	`bson:"sig" 			json:"sig"`
}

// type FundingTransaction struct {
// 	Transaction
// }

// type TriggerTransaction struct {
// 	Transaction
// }

// type SettlementTransaction struct {
// 	Transaction
// }

// type GeneralTransaction struct{
// 	FundingTransaction
// 	TriggerTransaction
// 	SettlementTransaction
// }

/*
Create :creating transactions
*/
func (t *Transaction) Create() (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.Flag = "Transaction"
	lib.Log.Debug("Create Transaction:", t)
	return true, nil
}

func (t *Transaction) CreateFundingTx(From int64, To string, Coin float32, Sig string) (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.Flag = "FundingTx"
	t.TimeLock = 0
	t.From = From
	t.To = To
	t.Coin = Coin
	t.NCommit = 0
	t.Sig = Sig
	lib.Log.Debug("Create Funding Transaction:", t)
	return true, nil
}

func (t *Transaction) CreateTriggerTx(From int64, To string, Coin float32, NCommit uint8, Sig string) (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.Flag = "TriggerTx"
	t.TimeLock = 100
	t.From = From
	t.To = To
	t.Coin = Coin
	t.NCommit = NCommit
	t.Sig = Sig
	lib.Log.Debug("Create Trigger Transaction:", t)
	return true, nil
}

func (t *Transaction) CreateSettlementTx(From int64, To string, Coin float32, NCommit uint8, Sig string) (bool, error) {
	t.ID, _ = lib.GetNewUID()
	t.Flag = "SettlementTx"
	t.From = From
	t.To = To
	t.Coin = Coin
	t.NCommit = NCommit
	t.Sig = Sig
	lib.Log.Debug("Create Settlement Transaction:", t)
	return true, nil
}
