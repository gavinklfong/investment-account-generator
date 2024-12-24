package investment

type Account struct {
	Number        string         `json:"accountNumber"`
	StockHoldings map[string]int `json:"stockHoldings"`
}

func NewAccount(number string) *Account {
	account := new(Account)
	account.Number = number
	account.StockHoldings = map[string]int{}
	return account
}
