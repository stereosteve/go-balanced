package balanced

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const secret = "d0342b8a10e611e3adaf026ba7d31e6f"

func TestNewClient(t *testing.T) {
	c := NewClient(nil, "")

	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, want %v", c.BaseURL.String(), defaultBaseURL)
	}
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, want %v", c.UserAgent, userAgent)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil, "")

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &Customer{Name: "l"}, `{"name":"l"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v) Body = %v, want %v", inBody, string(body), outBody)
	}

	// test that default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, want %v", userAgent, c.UserAgent)
	}
}

func TestCreateCustomer(t *testing.T) {
	// t.SkipNow()

	c := NewClient(nil, secret)
	u := "/v1/customers"
	inBody := &Customer{Name: "Go Balanced"}

	req, _ := c.NewRequest("POST", u, inBody)
	cust := new(Customer)
	resp, err := c.Do(req, cust)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 201 {
		t.Errorf("Expected 201, got %v", resp.StatusCode)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(cust)
}

func TestListCustomers(t *testing.T) {
	c := NewClient(nil, secret)

	customers, err := c.Customers.List()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(customers)
}
