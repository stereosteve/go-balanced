package balanced

import (
	"fmt"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	c := NewClient(nil, secret)
	inBody := &Customer{Name: "Go Balanced"}

	cust, err := c.Customers.Create(inBody)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cust)
}

func TestListCustomers(t *testing.T) {
	c := NewClient(nil, secret)

	opts := &ListOptions{Limit: 2, Offset: 1, Uri: "/v1/customers?limit=1"}
	customers, page, err := c.Customers.List(opts)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(page.Total)
	fmt.Println(customers)
}
