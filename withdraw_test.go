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
		Valor:  "100",
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

func TestWithdrawList(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.WithdrawService().ListWithdraw(&ras.WithdrawListQuery{Filtro: 0})
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

func TestWithdrawListPending(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.WithdrawService().ListWithdraw(&ras.WithdrawListQuery{Filtro: 1})
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

func TestWithdrawListIDQuery(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.WithdrawService().ListWithdraw(&ras.WithdrawListQuery{Filtro: 0, ID: os.Getenv("WITHDRAWID_FILTER_TEST")})
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
