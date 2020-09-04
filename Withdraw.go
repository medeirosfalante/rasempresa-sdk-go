package ras

// WithdrawService struct
type WithdrawService struct {
	client *APIClient
}

//WithdrawCreate - struct para criar uma retirada
type WithdrawCreate struct {
	Taxa   int    `json:"taxa"`
	Moeda  string `json:"moeda"`
	Wallet string `json:"wallet"`
	Valor  string `json:"valor"`
}

//WithdrawResponse - struct de resposta ao criar uma subconta
type WithdrawResponse struct {
	COD      string `json:"COD"`
	ID       string `json:"ID"`
	Mensagem string `json:"Mensagem"`
}

// FeeResponse - retorno da taxa
type FeeResponse struct {
	TAXAS map[string]FeeItemResponse `json:"TAXAS"`
}

// FeeItemResponse - retorno
type FeeItemResponse struct {
	Slow    string `json:"SLOW"`
	Average string `json:"AVERAGE"`
	Fast    string `json:"FAST"`
}

// CreateWithdraw - criar um saque
func (p WithdrawService) CreateWithdraw(withdraw *WithdrawCreate) (*WithdrawResponse, *Error, error) {
	var response *WithdrawResponse
	err, errAPI := p.client.Request("POST", "/v1/privado/retirada_empresa", withdraw, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//WithdrawService - Instanciar serviço de saque
func (c *APIClient) WithdrawService() *WithdrawService {
	return &WithdrawService{client: c}
}

// Fees - lista o fee
func (p WithdrawService) Fees() (*FeeResponse, *Error, error) {
	var response *FeeResponse
	err, errAPI := p.client.Request("POST", "/v1/privado/taxas_retirada", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
