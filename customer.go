package balanced

/*
Customer represents a Balanced Customer
*/
type Customer struct {
	Id    string `json:"id,omitempty"`
	Uri   string `json:"uri,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type CustomerService struct {
	client *Client
}

func (s *CustomerService) List() ([]Customer, error) {
	u := "/v1/customers"

	page := new(Page)
	req, _ := s.client.NewRequest("GET", u, nil)
	_, err := s.client.Do(req, page)

	if err != nil {
		return nil, err
	}

	customers := []Customer{}
	err = page.CastItems(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
