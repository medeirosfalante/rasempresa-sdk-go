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
