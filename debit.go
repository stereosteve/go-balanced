package balanced

import "fmt"

/*
Debit
*/
type Debit struct {
	OnBehalfOfUri string `json:"on_behalf_of_uri,omitempty"`
	Amount        int    `json:"amount,omitempty"`
	Description   string `json:"description,omitempty"`
	Uri           string `json:"uri,omitempty"`
}

type DebitService struct {
	client *Client
}

func (s *DebitService) Create(customer *Customer, debit *Debit) (*Debit, error) {
	u := fmt.Sprintf("%s/debits", customer.Uri)
	req, err := s.client.NewRequest("POST", u, debit)
	if err != nil {
		return nil, err
	}
	d := new(Debit)
	_, err = s.client.Do(req, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
