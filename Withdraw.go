package ras

// WithdrawService struct
type WithdrawService struct {
	client *APIClient
}

//WithdrawService - Instanciar serviço de saque
func (c *APIClient) WithdrawService() *WithdrawService {
	return &WithdrawService{client: c}
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

// WithdrawListQuery - query for list \n  0 =  todos, 1 =  em analise, 2  = aprovados, 3  = reprovados
type WithdrawListQuery struct {
	Filtro int    `json:"filtro"`
	ID     string `json:"id_retirada"`
}

//WithdrawItem - item do saque
type WithdrawItem struct {
	ID            string `json:"id_retirada"`
	Data          string `json:"data"`
	Hora          string `json:"hora"`
	Moeda         string `json:"moeda"`
	Valor         string `json:"valor"`
	Taxa          string `json:"taxa"`
	ValorRecebido string `json:"valor_recebido"`
	Comprovante   string `json:"comprovante"`
	Wallet        string `json:"wallet"`
	Banco         string `json:"banco"`
	Conta         string `json:"conta"`
	Agencia       string `json:"agencia"`
	Status        string `json:"status"`
}

// WithdrawListResponse - Retorno da lista de retiradas
type WithdrawListResponse struct {
	COD      string          `json:"COD"`
	Data     []*WithdrawItem `json:"data"`
	Mensagem string          `json:"Mensagem"`
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

// Fees - lista o fee
func (p WithdrawService) ListWithdraw(query *WithdrawListQuery) (*WithdrawListResponse, *Error, error) {
	var response *WithdrawListResponse
	err, errAPI := p.client.Request("POST", "/v1/privado/listar_retiradas_empresa", query, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
