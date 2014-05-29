package balanced

import "fmt"

/*
Credit
*/
type Credit struct {
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Uri         string `json:"uri,omitempty"`
}

type CreditService struct {
	client *Client
}

func (s *CreditService) CreateForCustomer(customer *Customer, credit *Credit) (*Credit, error) {
	u := fmt.Sprintf("%s/credits", customer.Uri)
	req, err := s.client.NewRequest("POST", u, credit)
	if err != nil {
		return nil, err
	}
	c := new(Credit)
	_, err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
