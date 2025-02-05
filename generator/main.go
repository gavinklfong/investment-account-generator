package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"example.com/investment"
	"example.com/investment/encoding"
)

const (
	AccountPrefix  = "INV"
	BatchCount     = 100
	BatchSize      = 10000
	OutputPath     = "./output"
	MaxHoldingUnit = 100
	StartDate      = "2024-01-01"
	EndDate        = "2024-12-31"
	Encoding       = "CSV"
)

var TickerList = []string{"AAPL", "SBUX", "MSFT", "CSCO", "QCOM", "META", "AMZN", "TSLA", "AMD", "NFLX"}

func main() {
	defer timer("main")()
	log.SetFlags(0)

	if yes, _ := exists(OutputPath); !yes {
		os.Mkdir(OutputPath, os.ModePerm)
	}

	c := make(chan struct{})

	for batch := 0; batch < BatchCount; batch++ {
		start := batch*BatchSize + 1
		end := start + BatchSize - 1
		go func() {
			generateAccountHoldingWithBatchSeq(batch, start, end)
			c <- struct{}{}
		}()
	}

	for batch := 0; batch < BatchCount; batch++ {
		<-c
	}

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func generateAccountHolding(suffix int, date time.Time) *investment.AccountHolding {
	account := investment.NewAccountHolding(fmt.Sprintf("%v-%010d", AccountPrefix, suffix), date)
	tickerCount := rand.Intn(len(TickerList))
	for _, value := range rand.Perm(tickerCount) {
		account.StockHoldings[TickerList[value]] = rand.Intn(MaxHoldingUnit)
	}
	return account
}

func generateAccountHoldingWithBatchSeq(batch, start, end int) {

	file, err := os.OpenFile(fmt.Sprintf("%v/investment-account-%01d.%v", OutputPath, batch, strings.ToLower(Encoding)), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	writer := encoding.NewWriter(Encoding, TickerList, file)
	defer writer.End()

	if err = writer.Init(); err != nil {
		log.Panic(err)
	}

	for seq := start; seq <= end; seq++ {
		generateAccountHoldingWithDateRange(writer, seq)
	}
}

func generateAccountHoldingWithDateRange(writer encoding.Writer, accountSeq int) {

	startDate, err := time.Parse(time.DateOnly, StartDate)
	if err != nil {
		log.Fatalf("fail to parse start date: %v", StartDate)
	}

	endDate, err := time.Parse(time.DateOnly, EndDate)
	if err != nil {
		log.Fatalf("fail to parse end date: %v", EndDate)
	}

	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		accountHolding := generateAccountHolding(accountSeq, date)
		if err = writer.Write(accountHolding); err != nil {
			log.Panic(err)
		}
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
