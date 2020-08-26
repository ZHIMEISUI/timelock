
package lib

import(
	"strings"
)

func TxHandle(tx string) map[string]string {
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