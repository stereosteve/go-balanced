package balanced

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// A BalancedError is returned when the response has a
// 4xx or 5xx status code.
type BalancedError struct {
	Status       string `json:"status"`
	CategoryCode string `json:"category_code"`
	CategoryType string `json:"category_type"`
	Description  string `json:"description"`
	RequestId    string `json:"request_id"`
	StatusCode   int    `json:"status_code"`
}

func (b *BalancedError) Error() string {
	return fmt.Sprintf("%d: %s - %s", b.StatusCode, b.Status, b.Description)
}

// Client
type Client struct {
	Secret string
}

// Url - constructs URL with auth from a partial path (e.g. /v1/customers)
func (c *Client) Url(path string) string {
	parts := []string{"https://", c.Secret, ":@api.balancedpayments.com", path}
	return strings.Join(parts, "")
}

// Get
func (c *Client) Get(path string, v interface{}) error {
	var args interface{}
	return c.Do("GET", path, args, v)
}

// Post
func (c *Client) Post(path string, args interface{}, v interface{}) error {
	return c.Do("POST", path, args, v)
}

// Put
func (c *Client) Put(path string, args interface{}, v interface{}) error {
	return c.Do("PUT", path, args, v)
}

// Delete
func (c *Client) Delete(path string) error {
	req, err := http.NewRequest("DELETE", c.Url(path), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return errors.New(resp.Status)
	}
	return nil
}

// Do
func (c *Client) Do(method string, path string, args interface{}, v interface{}) error {

	// encode json
	body, err := json.Marshal(args)
	if err != nil {
		return err
	}

	// construct request
	req, err := http.NewRequest(method, c.Url(path), bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)

	// handle non-2xx response
	if resp.StatusCode >= 400 {
		balancedError := BalancedError{}
		decoder.Decode(&balancedError)
		return &balancedError
	}

	// decode response
	if err = decoder.Decode(&v); err != nil {
		return err
	}

	return nil
}

type Customer struct {
	Id    string
	Uri   string
	Name  string
	Phone string
}

type BankAccount struct {
	Uri         string
	Credits_uri string
	Bank_name   string
	Can_debit   bool
}

type Debit struct {
	Uri                string
	Status             string
	Transaction_number string
}
