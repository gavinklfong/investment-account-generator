package investment

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	accountNumber := "ABC0001"
	account := NewAccount(accountNumber)
	if account.Number != accountNumber {
		t.Fatalf("NewAccount number does not match")
	}
}
