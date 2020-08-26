package lib


func has(strs []string, str string, index string) (string, bool) {
	for _,t := range strs{
		txmap := txHandle(t)
		if str == txmap[index] {
			return t, true
		}
	}
	return "",false
}