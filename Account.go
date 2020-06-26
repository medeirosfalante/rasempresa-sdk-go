package ras

import (
	"errors"
)

// AccountService struct
type AccountService struct {
	client *APIClient
}

//SubAccountCreate - struct para criar uma subconta
type SubAccountCreate struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

type BalanceAccountRequest struct {
	Subconta string `json:"subconta"`
}

type BalanceAccountResponse struct {
	COD      string       `json:"COD"`
	Mensagem string       `json:"Mensagem"`
	Cadastro *AccountData `json:"cadastro"`
	Wallets  wallets      `json:"wallets"`
	Saldos   saldos       `json:"saldos"`
}

type ListSubAccountsResponse struct {
	COD      string           `json:"COD"`
	Mensagem string           `json:"Mensagem"`
	Cadastro []SubAccountItem `json:"data"`
}

type SubAccountItem struct {
	IDSubconta string `json:"id_subconta"`
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	SaldoRAS   string `json:"saldo_ras"`
	SaldoBTC   string `json:"saldo_btc"`
	SaldoETH   string `json:"saldo_eth"`
	SaldoDASH  string `json:"saldo_dash"`
	SaldoUSDT  string `json:"saldo_usdt"`
	SaldoBRL   string `json:"saldo_brl"`
	SaldoDOLAR string `json:"saldo_dolar"`
	Status     string `json:"status"`
}

type AccountData struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Data     string `json:"data"`
	Telefone string `json:"telefone"`
	Cpf      string `json:"cpf"`
	Rg       string `json:"rg"`
	Banco    string `json:"banco"`
	Ag       string `json:"ag"`
	Conta    string `json:"conta"`
}

type wallets map[string]string
type saldos map[string]string

//SubAccountResponse - struct de resposta ao criar uma subconta
type SubAccountResponse struct {
	COD      string `json:"COD"`
	ID       string `json:"ID"`
	Mensagem string `json:"Mensagem"`
}

//AccountService - gestão das contas
func (c *APIClient) AccountService() *AccountService {
	return &AccountService{client: c}
}

// CreateSubAccount - criar uma subconta
func (p AccountService) CreateSubAccount(subaccount *SubAccountCreate) (*SubAccountResponse, *Error, error) {
	var response *SubAccountResponse
	if subaccount == nil {
		return nil, nil, errors.New("subaccount is null")
	}
	err, errAPI := p.client.Request("POST", "/v1/privado/criar_subconta", subaccount, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//BalanceSubAccount - detalhes e saldo da subconta
func (p AccountService) BalanceSubAccount(subconta string) (*BalanceAccountResponse, *Error, error) {
	var response *BalanceAccountResponse
	if subconta == "" {
		return nil, nil, errors.New("subconta is empty")
	}
	err, errAPI := p.client.Request("POST", "/v1/privado/exibir_balanco_subconta", &BalanceAccountRequest{Subconta: subconta}, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//ListSubAccount - listar subcontas
func (p AccountService) ListSubAccount() (*ListSubAccountsResponse, *Error, error) {
	var response *ListSubAccountsResponse
	err, errAPI := p.client.Request("GET", "/v1/privado/listar_subconta", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
