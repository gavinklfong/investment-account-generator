package investment

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
