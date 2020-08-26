package lib


func logTx(funcname string, txmap map[string]string){
	Log.Debug(funcname+" starts Debug...")
	Log.Debug("Transaction ID: "+txmap["ID"])
	Log.Debug("Transaction Type: "+txmap["Flag"])
	Log.Debug("TimeLock: "+txmap["TimeLock"])
	Log.Debug("From: "+txmap["From"])
	Log.Debug("To: "+txmap["To"])
	Log.Debug("Deposit Coins: "+txmap["Coin"])
	Log.Debug("Channel Version: "+txmap["NCommit"])
	Log.Debug("Secret T: "+txmap["SecretT"])
	Log.Debug("Sig: "+txmap["Sig"])
	Log.Debug(funcname+" ends Debug...")
}