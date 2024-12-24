package main

import (
	// "encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"

	"example.com/investment"
	"example.com/investment/encoding/csv"
)

const (
	AccountPrefix  = "INV"
	BatchSize      = 10
	OutputPath     = "./output"
	MaxHoldingUnit = 100
)

var TickerList = []string{"AAPL", "SBUX", "MSFT", "CSCO", "QCOM", "META", "AMZN", "TSLA", "AMD", "NFLX"}

type AccountWriter interface {
	Write(account *investment.Account) string
}

func Insert(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}

func main() {

	log.SetFlags(0)

	for batch := 0; batch < 2; batch++ {
		start := batch*BatchSize + 1
		end := start + BatchSize - 1
		generateAndWriteAccount(batch, start, end)
	}

}

func generateAccount(suffix int) *investment.Account {
	account := investment.NewAccount(fmt.Sprintf("%v-%010d", AccountPrefix, suffix))
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

	writer := csv.NewWriter(TickerList, file)
	defer writer.Flush()

	if err = writer.Init(); err != nil {
		log.Panic(err)
	}

	for seq := start; seq <= end; seq++ {
		account := generateAccount(seq)
		if err = writer.Write(account); err != nil {
			log.Panic(err)
		}
	}
}
