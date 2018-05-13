package lib

import (
	"encoding/json"

	"../gdax"
)

func UnmarshalProducts(rawProducts []byte) map[string]gdax.Product {
	productList := []gdax.Product{}
	err := json.Unmarshal(rawProducts, &productList)
	if err != nil {
		panic(err)
	}

	l := len(productList)
	products := make(map[string]gdax.Product)

	for i := 0; i < l; i++ {
		products[productList[i].ID] = productList[i]
	}

	return products
}

func UnmarshalTicker(rawTicker []byte) gdax.Ticker {
	ticker := gdax.Ticker{}

	err := json.Unmarshal(rawTicker, &ticker)
	if err != nil {
		panic(err)
	}

	return ticker
}
