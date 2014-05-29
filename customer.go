package balanced

import (
	"net/url"
	"strconv"
)

/*
Customer represents a Balanced Customer
*/
type Customer struct {
	Id    string `json:"id,omitempty"`
	Uri   string `json:"uri,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type CustomerService struct {
	client *Client
}

func (s *CustomerService) List(opt *ListOptions) ([]Customer, *Page, error) {
	u := "/v1/customers"

	if opt != nil {
		params := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
		}
		u += "?" + params.Encode()
		if opt.Uri != "" {
			u = opt.Uri
		}
	}

	page := new(Page)
	req, _ := s.client.NewRequest("GET", u, nil)
	_, err := s.client.Do(req, page)

	if err != nil {
		return nil, nil, err
	}

	customers := []Customer{}
	err = page.CastItems(&customers)
	if err != nil {
		return nil, nil, err
	}

	return customers, page, nil
}

func (s *CustomerService) Create(customer *Customer) (*Customer, error) {
	u := "/v1/customers"
	req, err := s.client.NewRequest("POST", u, customer)
	if err != nil {
		return nil, err
	}
	c := new(Customer)
	_, err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CustomerService) AddBankAccount(customer *Customer, bankUri string) (*Customer, error) {
	data := struct {
		BankAccountUri string `json:"bank_account_uri"`
	}{bankUri}
	req, err := s.client.NewRequest("PUT", customer.Uri, data)
	if err != nil {
		return nil, err
	}
	c := new(Customer)
	_, err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CustomerService) AddCard(customer *Customer, cardUri string) (*Customer, error) {
	data := struct {
		CardUri string `json:"card_uri"`
	}{cardUri}
	req, err := s.client.NewRequest("PUT", customer.Uri, data)
	if err != nil {
		return nil, err
	}
	c := new(Customer)
	_, err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
