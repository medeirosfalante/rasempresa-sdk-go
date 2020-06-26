package ras_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"
	ras "github.com/rafaeltokyo/rasempresa-sdk-go"
)

func TestCreateAccount(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.AccountService().CreateSubAccount(&ras.SubAccountCreate{Nome: "Teste"})
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

func TestBalanceSubAccount(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.AccountService().BalanceSubAccount(os.Getenv("SUBACCOUNT_TEST"))
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

func TestBalanceSubAccountValidAllAddress(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.AccountService().BalanceSubAccount(os.Getenv("SUBACCOUNT_TEST"))
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
	re := regexp.MustCompile("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$")
	if !re.MatchString(response.Wallets["BTC"]) {
		t.Error("invalid BTC address")
		return

	}
	re = regexp.MustCompile("^([X7])[a-zA-Z0-9]{33}$")
	if !re.MatchString(response.Wallets["DASH"]) {
		t.Error("invalid DASH address")
		return
	}
	re = regexp.MustCompile("^0x[a-fA-F0-9]{40}$")
	if !re.MatchString(response.Wallets["RAS"]) {
		t.Error("invalid RAS address")
		return

	}
	re = regexp.MustCompile("^0x[a-fA-F0-9]{40}$")
	if !re.MatchString(response.Wallets["USDT"]) {
		t.Error("invalid USDT address")
		return

	}
}

func TestListSubAccount(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.AccountService().ListSubAccount()
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
