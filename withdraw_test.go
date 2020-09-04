package ras_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	ras "github.com/rafaeltokyo/rasempresa-sdk-go"
)

func TestGetFee(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.WithdrawService().Fees()
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestWithdraw(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.WithdrawService().CreateWithdraw(&ras.WithdrawCreate{
		Taxa:   2,
		Moeda:  "RAS",
		Wallet: os.Getenv("WALLET_RAS_WITHDRAW"),
		Valor:  "10",
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}
