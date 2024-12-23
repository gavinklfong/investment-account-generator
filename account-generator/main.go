package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

const AccountPrefix = "INV"
const BatchSize = 10
const OutputPath = "./output"
const MaxHoldingUnit = 100

var TickerList = []string{"AAPL", "SBUX", "MSFT", "CSCO", "QCOM", "META", "AMZN", "TSLA", "AMD", "NFLX"}

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

func main2() {
	log.SetFlags(0)
	generateAccount(1)
}

func generateAccount(suffix int) *Account {
	account := NewAccount(fmt.Sprintf("%v-%010d", AccountPrefix, suffix))
	tickerCount := rand.Intn(len(TickerList))
	for _, value := range rand.Perm(tickerCount) {
		account.StockHoldings[TickerList[value]] = rand.Intn(MaxHoldingUnit)
	}
	// fmt.Printf("%v\n", *account)
	return account
}

func generateAndWriteAccount(batch, start, end int) {

	f, err := os.OpenFile(fmt.Sprintf("%v/investment-account-%01d.csv", OutputPath, batch), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Panic(err)
	}

	for seq := start; seq <= end; seq++ {
		// account := NewAccount(fmt.Sprintf("%v-%010d", AccountPrefix, seq))
		account := generateAccount(seq)
		fmt.Printf("%v\n", *account)
		if _, err = io.WriteString(f, fmt.Sprintln(account.Number)); err != nil {
			log.Panic(err)
		}
	}
}

func writeFileExample() {
	f, err := os.OpenFile("./test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("testing\n"); err != nil {
		panic(err)
	}

	if _, err = io.WriteString(f, "testing (written by io package)\n"); err != nil {
		panic(err)
	}

}
