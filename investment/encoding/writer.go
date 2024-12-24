package encoding

import (
	"example.com/investment"
)

type Writer interface {
	Init() error
	Write(account *investment.Account) error
	Flush()
}

// func NewWriter(encoding string, tickerList []string, w io.Writer) (writer *Writer) {
// 	switch encoding {
// 	case "CSV":
// 		writer = csv.NewWriter(tickerList, w)
// 	}
// }
