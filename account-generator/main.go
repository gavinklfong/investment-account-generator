package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	AccountPrefix  = "INV"
	BatchSize      = 10
	OutputPath     = "./output"
	MaxHoldingUnit = 100
)

var TickerList = []string{"AAPL", "SBUX", "MSFT", "CSCO", "QCOM", "META", "AMZN", "TSLA", "AMD", "NFLX"}

type AccountWriter interface {
	Write(account *Account) string
}

type JSONWriter struct{}
type CSVWriter struct {
	w *csv.Writer
}

func NewCSVWriter(w io.Writer) *CSVWriter {
	return &CSVWriter{
		w: csv.NewWriter(w),
	}
}

func (w *CSVWriter) WriteHeader() error {
	return w.w.Write(Insert(TickerList, "AccountNumber", 0))
}

func (w *CSVWriter) Write(account *Account) error {
	var fields [11]string
	fields[0] = account.Number

	for i, ticker := range TickerList {
		unit, ok := account.StockHoldings[ticker]
		if ok {
			fields[i+1] = strconv.Itoa(unit)
		} else {
			fields[i+1] = "0"
		}
	}

	fmt.Println(strings.Join(fields[:], ","))

	return w.w.Write(fields[:])
}

func (w *CSVWriter) Flush() {
	w.w.Flush()
}

func Insert(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}

type Account struct {
	Number        string
	StockHoldings map[string]int
}

func NewAccount(number string) *Account {
	account := new(Account)
	account.Number = number
	account.StockHoldings = map[string]int{}
	return account
}

func main() {

	log.SetFlags(0)

	for batch := 0; batch < 2; batch++ {
		start := batch*BatchSize + 1
		end := start + BatchSize - 1
		generateAndWriteAccount(batch, start, end)
	}

}

func generateAccount(suffix int) *Account {
	account := NewAccount(fmt.Sprintf("%v-%010d", AccountPrefix, suffix))
	tickerCount := rand.Intn(len(TickerList))
	for _, value := range rand.Perm(tickerCount) {
		account.StockHoldings[TickerList[value]] = rand.Intn(MaxHoldingUnit)
	}
	return account
}

func generateAndWriteAccount(batch, start, end int) {

	file, err := os.OpenFile(fmt.Sprintf("%v/investment-account-%01d.csv", OutputPath, batch), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	writer := NewCSVWriter(file)
	defer writer.Flush()

	if err = writer.WriteHeader(); err != nil {
		log.Panic(err)
	}

	for seq := start; seq <= end; seq++ {
		account := generateAccount(seq)
		if err = writer.Write(account); err != nil {
			log.Panic(err)
		}
	}
}
