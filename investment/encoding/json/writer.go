package json

import (
	"encoding/json"
	"io"

	"example.com/investment"
)

type Writer struct {
	TickerList []string
	w          io.Writer
}

func NewWriter(tickerList []string, w io.Writer) *Writer {
	return &Writer{
		TickerList: tickerList,
		w:          w,
	}
}

func (w *Writer) Init() error {
	_, err := w.w.Write([]byte("[\n"))
	return err
}

func (w *Writer) Write(accountHolding *investment.AccountHolding) error {
	b, err := json.Marshal(accountHolding)
	if err != nil {
		return err
	}

	_, err = w.w.Write(append(b, ",\n"...))
	return err
}

func (w *Writer) End() {
	w.w.Write([]byte("]"))
}
