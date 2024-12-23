package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const BatchSize = 1000000
const OutputPath = "./output"

type Account struct {
	Number        string
	StockHoldings map[string]float32
}

func NewAccount(number string) *Account {
	account := new(Account)
	account.Number = number
	account.StockHoldings = map[string]float32{}
	return account
}

func main() {

	log.SetFlags(0)

	for batch := 0; batch < 2; batch++ {
		start := batch*BatchSize + 1
		end := start + BatchSize - 1
		generateAccount(batch, start, end)
	}

}

func generateAccount(batch, start, end int) {

	f, err := os.OpenFile(fmt.Sprintf("%v/investment-account-%01d.csv", OutputPath, batch), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Panic(err)
	}

	for seq := start; seq <= end; seq++ {
		account := NewAccount(fmt.Sprintf("INV-%010d", seq))
		// log.Println(accountNumber)
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
