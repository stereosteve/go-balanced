package balanced

/*
BankAccount
*/
type BankAccount struct {
	Name          string `json:"name,omitempty"`
	RoutingNumber string `json:"routing_number,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	Type          string `json:"type,omitempty"`
	Uri           string `json:"uri,omitempty"`
	BankName      string `json:"bank_name,omitempty"`
	CanDebit      bool   `json:"can_debit,omitempty"`
}

type BankAccountService struct {
	client *Client
}

func (s *BankAccountService) Create(bank *BankAccount) (*BankAccount, error) {
	u := "/v1/bank_accounts"
	req, err := s.client.NewRequest("POST", u, bank)
	if err != nil {
		return nil, err
	}

	ba := new(BankAccount)
	_, err = s.client.Do(req, ba)
	if err != nil {
		return nil, err
	}
	return ba, nil
}
