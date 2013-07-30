package balanced

import (
	"fmt"
	"testing"
)

var c *Client
var customer *Customer

func TestUrl(t *testing.T) {
	c = &Client{"9a946c52e98011e282f9026ba7d31e6f"}
	in := "/path"
	out := c.Url(in)
	want := "https://9a946c52e98011e282f9026ba7d31e6f:@api.balancedpayments.com/path"
	if out != want {
		t.Errorf("Url(%s) = %s, \n want %s", in, out, want)
	}
}

func TestCreateCustomer(t *testing.T) {
	args := map[string]interface{}{
		"name":  "gopher",
		"phone": "123-456-7890",
	}
	err := c.Post("/v1/customers", &args, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestGetCustomer(t *testing.T) {
	var c2 map[string]interface{}
	err := c.Get(customer.Uri, &c2)
	if err != nil {
		panic(err)
	}
	if customer.Id != c2["id"] {
		t.Errorf("Customer ids not equal %s, %s", customer.Id, c2["id"])
	}
}

func TestUpdateCustomer(t *testing.T) {
	args := map[string]interface{}{
		"name": "Updated Gopher",
	}
	err := c.Put(customer.Uri, &args, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestDeleteCustomer(t *testing.T) {
	err := c.Delete(customer.Uri)
	if err != nil {
		panic(err)
	}
}
