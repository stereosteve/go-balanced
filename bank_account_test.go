package balanced

import (
	"fmt"
	"testing"
)

func TestCreateBankAccount(t *testing.T) {
	c := NewClient(nil, secret)
	inBank := &BankAccount{
		Name:          "Go Account",
		RoutingNumber: "021000021",
		AccountNumber: "9900000002",
		Type:          "checking",
	}
	outBank, err := c.BankAccounts.Create(inBank)
	if err != nil {
		t.Errorf("failed to create bank account: %v", err)
	}
	fmt.Println(outBank)
}
