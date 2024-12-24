package encoding

import (
	"io"
	"log"

	"example.com/investment"
	"example.com/investment/encoding/csv"
	"example.com/investment/encoding/json"
)

type Writer interface {
	Init() error
	Write(account *investment.Account) error
	End()
}

func NewWriter(encoding string, tickerList []string, w io.Writer) (writer Writer) {
	switch encoding {
	case "CSV":
		writer = csv.NewWriter(tickerList, w)
	case "JSON":
		writer = json.NewWriter(tickerList, w)
	default:
		log.Panic("unknown encoding type")
	}
	return
}
