package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

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

func (w *Writer) Write(account *investment.Account) error {
	var fields [11]string
	fields[0] = account.Number

	for i, ticker := range w.TickerList {
		unit, ok := account.StockHoldings[ticker]
		if ok {
			fields[i+1] = strconv.Itoa(unit)
		}
	}

	fmt.Println(strings.Join(fields[:], ","))

	return w.w.Write(fields[:])
}

func (w *Writer) End() {
	w.w.Flush()
}

func insert(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}
