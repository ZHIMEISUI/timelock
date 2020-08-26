package lib


func Has(strs []string, str string, index string) (string, bool) {
	for _,t := range strs{
		txmap := TxHandle(t)
		if str == txmap[index] {
			return t, true
		}
	}
	return "",false
}