package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Account struct {
	Number string
}

func main() {

	log.SetFlags(0)

	for batch := 0; batch < 5; batch++ {
		start := batch*1000000 + 1
		end := start + 1000000 - 1
		generateAccount(batch, start, end)
	}

}

func generateAccount(batch, start, end int) {

	f, err := os.OpenFile(fmt.Sprintf("./investment-account-%01d.csv", batch), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Panic(err)
	}

	for seq := start; seq <= end; seq++ {
		accountNumber := fmt.Sprintf("INV-%010d", seq)
		// log.Println(accountNumber)
		if _, err = io.WriteString(f, fmt.Sprintln(accountNumber)); err != nil {
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
