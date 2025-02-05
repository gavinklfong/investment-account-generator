package investment

import (
	"testing"
	"time"
)

func TestNewAccountHolding(t *testing.T) {
	accountNumber := "ABC0001"
	currentDate := time.Now().Truncate(24 * time.Hour)
	accountHolding := NewAccountHolding(accountNumber, currentDate)
	if accountHolding.Number != accountNumber {
		t.Fatalf("NewAccountHolding number does not match")
	}
	if accountHolding.Date != currentDate {
		t.Fatalf("NewAccountHolding date does not match")
	}
}
