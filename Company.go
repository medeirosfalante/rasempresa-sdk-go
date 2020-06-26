package ras

//CompanyService - serviço de empresa
type CompanyService struct {
	client *APIClient
}

// CompanyInfo - informações sobre seus dados
type CompanyInfo struct {
	COD          string `json:"COD"`
	Mensagem     string `json:"Mensagem"`
	ID           string `json:"id"`
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	Data         string `json:"data"`
	Telefone     string `json:"telefone"`
	Cpf          string `json:"cpf"`
	Rg           string `json:"rg"`
	Banco        string `json:"banco"`
	Agencia      string `json:"agencia"`
	Conta        string `json:"conta"`
	RazaoSocial  string `json:"razao_social"`
	TaxaComissao string `json:"taxa_comissao"`
}

//CompanyService - Instanciar serviço de empresa
func (c *APIClient) CompanyService() *CompanyService {
	return &CompanyService{client: c}
}

//GetInfo - detalhes da conta
func (p CompanyService) GetInfo() (*CompanyInfo, *Error, error) {
	var response *CompanyInfo
	err, errAPI := p.client.Request("GET", "/v1/privado/exibir_dados_empresa", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
