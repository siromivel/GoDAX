package gdax

// Product represents a product available through the GDAX REST API
type Product struct {
	ID             string
	BaseCurrency   string
	BaseMinSize    float32
	BaseMaxSize    float32
	QuoteCurrency  string
	QuoteIncrement float32
}
