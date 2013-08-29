package balanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://api.balancedpayments.com/"
	userAgent      = "go-balanced/" + libraryVersion
)

// Client
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Secret    string
}

// NewClient returns a new Balanced API client.  If nil httpClient is provided
// http.DefaultClient will be used.  A Balanced Secret must be provided.
func NewClient(httpClient *http.Client, secret string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Secret: secret}
	return c
}

// NewRequest creates a Balanced API request, resolvig a relative URL into a
// full URL, relative the client's BaseURL, and sending the client.Secret.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.Secret, "")
	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

/*
An ErrorResponse reports a Balanced error
*/
type ErrorResponse struct {
	Response     *http.Response // HTTP response that caused this error
	Status       string         `json:"status"`
	CategoryCode string         `json:"category_code"`
	CategoryType string         `json:"category_type"`
	Description  string         `json:"description"`
	RequestId    string         `json:"request_id"`
	StatusCode   int            `json:"status_code"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Description)
}

// CheckResponse checks a response for errors and returns an ErrorResponse
// if the response StatusCode is outside of the valid 2xx range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

/*
Page has items and links to other pages
*/
type Page struct {
	Uri     string          `json:"uri,omitempty"`
	NextURI string          `json:"next_uri",omitempty"`
	LastURI string          `json:"last_uri",omitempty"`
	Items   json.RawMessage `json:"items",omitempty"`
}

// casts Items RawMessage to a desired type
func (p *Page) CastItems(v interface{}) error {
	err := json.Unmarshal(p.Items, &v)
	return err
}

/*
Customer represents a Balanced Customer
*/
type Customer struct {
	Id    string `json:"id,omitempty"`
	Uri   string `json:"uri,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

/*
BankAccount
*/
type BankAccount struct {
	Uri         string
	Credits_uri string
	Bank_name   string
	Can_debit   bool
}

/*
Debit
*/
type Debit struct {
	Uri                string
	Status             string
	Transaction_number string
}
