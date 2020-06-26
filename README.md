# rasempresa-sdk-go
Sdk api rasempresa



Exemplo de criação de subconta


```go
client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
response, errAPI, err := client.AccountService().CreateSubAccount(&ras.SubAccountCreate{Nome: "Teste"})
if err != nil {
 log.Fatalf("err : %s", err)
 return
}
if errAPI != nil {
 log.Fatalf("errAPI : %#v", errAPI)
 return
}
if response == nil {
 log.Fatalf("response is null")
 return
}
fmt.Printf("myid:%s",response.ID)
```


Exemplo pegar endereço de Bitcoin de subconta


```go
client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
response, errAPI, err := client.AccountService().BalanceSubAccount("id da sua subcointa")
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
fmt.Printf("endereço BTC", response.Wallets["BTC"])
```
