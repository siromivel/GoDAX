package lib

import (
	"encoding/json"
	"testing"

	"../gdax"
)

func TestUnmarshalProducts(t *testing.T) {
	expectedProducts := make([]gdax.Product, 2)
	expectedProducts[0] = gdax.Product{"ETH-USD", "ETH", 0.0001, 1000000.00, "USD", 0.01}
	expectedProducts[1] = gdax.Product{"ETH-BTC", "ETH", 0.0001, 1000000.00, "BTC", 0.00001}

	productBytes, _ := json.Marshal(expectedProducts)
	actualProducts := UnmarshalProducts(productBytes)

	if actualProducts["ETH-USD"] != expectedProducts[0] {
		t.Errorf("Error mapping products, expected: %+v, got: %+v", expectedProducts[0], actualProducts["ETH-USD"])
	}

	if actualProducts["ETH-BTC"] != expectedProducts[1] {
		t.Errorf("Error mapping products, expected: %+v, got: %+v", expectedProducts[1], actualProducts["ETH-BTC"])
	}
}

func TestUnmarshalTicker(t *testing.T) {
	expectedTicker := gdax.Ticker{420, 12.00, 1, 10.17, 12.13, 1000, "time4fun"}
	tickerBytes, _ := json.Marshal(expectedTicker)
	actualTicker := UnmarshalTicker(tickerBytes)

	if actualTicker != expectedTicker {
		t.Errorf("Error mapping ticker, expected: %+v, got: %+v", expectedTicker, actualTicker)
	}
}
