package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Market struct {
	Success string
	Message string
	Result  []struct {
		MarketName         string
		MarketCurrency     string
		BaseCurrency       string
		MarketCurrencyLong string
		BaseCurrencyLong   string
		IsActive           string
		MinTradeSize       float64
	}
}

type Balance struct {
	Success string
	Message string
	Result  []struct {
		Currency      string
		Balance       string
		Available     string
		Pending       string
		CryptoAddress string
		IsActive      string
		AllowDeposit  string
		AllowWithdraw string
	}
}

func main() {
	//url := "https://bleutrade.com/api/v2/public/getmarkets"
	//AuthAPI := "https://bleutrade.com/api/v2/account/getbalances"
	apiKey := "65e3698cb8f3129fe1254c4bff1e2f1b"
	apiSecret := "4bacffd6943a93745ce66e729ddcce5bc3c16036"
	nonce := time.Now().UnixNano()
	//fmt.Println(nonce)
	//nonce = nonce.String.Format("2006-01-02 15:04:05")
	url := fmt.Sprintf("https://bleutrade.com/api/v2/account/getbalances?apikey=%s&nonce=%d", apiKey, nonce)
	fmt.Println("Link antes de assinar:")
	fmt.Println(url)

	h := hmac.New(sha512.New, []byte(apiSecret))
	h.Write([]byte(url))
	sinal := hex.EncodeToString(h.Sum(nil))
	urlassinado := fmt.Sprintf(url+"&apisign=%s", sinal)
	fmt.Println("Url Assinado")
	fmt.Println(urlassinado)
	//defer resp.Body.Close()

	response, erro := http.Get(urlassinado)

	//if erro != nil {
	//	panic(erro)

	// lendo o json do response do http request
	responseJSON, erro := ioutil.ReadAll(response.Body)

	if erro != nil {
		panic(erro)
	}
	var dadosJSON Balance

	//convertendo dadosJSON
	erro = json.Unmarshal(responseJSON, &dadosJSON)

	if erro != nil {
		panic(erro)
	}
	fmt.Println("Saldos")
	fmt.Println(dadosJSON)

}
