package investment

import "time"

type AccountHolding struct {
	Number        string         `json:"accountNumber"`
	Date          time.Time      `json:"date"`
	StockHoldings map[string]int `json:"stockHoldings"`
}

func NewAccountHolding(number string, date time.Time) *AccountHolding {
	accountHolding := new(AccountHolding)
	accountHolding.Number = number
	accountHolding.Date = date
	accountHolding.StockHoldings = map[string]int{}
	return accountHolding
}
