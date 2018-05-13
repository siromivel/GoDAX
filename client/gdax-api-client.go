package client

import (
	"io/ioutil"
	"net/http"

	"../internal/gdax"
	"../internal/lib"
)

// GetProducts gets list of products from GDAX
func GetProducts() map[string]gdax.Product {
	response := sendGet("products")
	return lib.UnmarshalProducts(response)
}

// GetTicker gets ticker data for a product symbol
func GetTicker(symbol string) gdax.Ticker {
	response := sendGet("products/" + symbol + "/ticker")
	return lib.UnmarshalTicker(response)
}

func sendGet(target string) []byte {
	baseURL := "https://api.gdax.com/"
	res, err := http.Get(baseURL + target)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}
