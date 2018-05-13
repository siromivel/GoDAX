package gdax

// Ticker represents most recent trade and spread data about a product
type Ticker struct {
	LastTradeID uint64  `json:"trade_id"`
	Price       float64 `json:"price,string,omitempty"`
	Size        float64 `json:"size,string,omitempty"`
	Bid         float32 `json:"bid,string,omitempty"`
	Ask         float32 `json:"ask,string,omitempty"`
	Volume      float64 `json:"volume,string,omitempty"`
	Time        string
}
