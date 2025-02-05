package csv

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"example.com/investment"
)

type Writer struct {
	TickerList []string
	w          *csv.Writer
}

func NewWriter(tickerList []string, w io.Writer) *Writer {
	return &Writer{
		TickerList: tickerList,
		w:          csv.NewWriter(w),
	}
}

func (w *Writer) Init() error {
	return w.w.Write(insert(w.TickerList, "AccountNumber", 0))
}

func (w *Writer) Write(accountHolding *investment.AccountHolding) error {
	var fields [12]string
	fields[0] = accountHolding.Number
	fields[1] = accountHolding.Date.In(time.UTC).Format(time.DateOnly)

	for i, ticker := range w.TickerList {
		unit, ok := accountHolding.StockHoldings[ticker]
		if ok {
			fields[i+2] = strconv.Itoa(unit)
		}
	}

	return w.w.Write(fields[:])
}

func (w *Writer) End() {
	w.w.Flush()
}

func insert(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}
