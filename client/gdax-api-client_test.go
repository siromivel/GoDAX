package client

import (
	"reflect"
	"testing"

	"../internal/gdax"
)

func TestGetProducts(t *testing.T) {
	products := GetProducts()
	actualType := reflect.TypeOf(products["ETH-USD"])
	expectedType := reflect.TypeOf(gdax.Product{})

	if actualType != expectedType {
		t.Errorf("Error getting products, expected: %s, got: %s", expectedType.String(), actualType.String())
	}
}

func TestGetTicker(t *testing.T) {
	ticker := GetTicker("ETH-USD")
	actualType := reflect.TypeOf(ticker)
	expectedType := reflect.TypeOf(gdax.Ticker{})

	if actualType != expectedType {
		t.Errorf("Error getting products, expected: %s, got: %s", expectedType.String(), actualType.String())
	}
}
