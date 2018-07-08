package main

import (
	"encoding/json"
	"fmt"

	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-gdax"
)

type channelHandler func(*ws.Conn, chan string)

func main() {
	var tickerChan = make(chan string)

	products := []string{"ETH-USD", "ETH-BTC"}
	wsConn := openSocketConnection("wss://ws-feed.pro.coinbase.com")

	subscription := openSubscription(wsConn, "ticker", products)

	go channelStream(wsConn, subscription, tickerChan, tickerHandler)
	go channelPrinter(tickerChan)

	var input string
	fmt.Scanln(&input)
}

func openSocketConnection(endpoint string) *ws.Conn {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial(endpoint, nil)
	if err != nil {
		println(err.Error())
	}

	return wsConn
}

func openSubscription(wsConn *ws.Conn, channel string, products []string) gdax.Message {
	subscription := gdax.Message{
		Type: "subscribe",
		Channels: []gdax.MessageChannel{
			gdax.MessageChannel{
				Name:       channel,
				ProductIds: products,
			},
		},
	}
	if err := wsConn.WriteJSON(subscription); err != nil {
		println(err.Error())
	}

	return subscription
}

func channelStream(wsConn *ws.Conn, subscription gdax.Message, channel chan string, handler channelHandler) {
	message := gdax.Message{}
	if err := wsConn.ReadJSON(&message); err != nil {
		println(err.Error())
	}
	handler(wsConn, channel)
}

func tickerHandler(wsConn *ws.Conn, tickerChan chan string) {
	for true {
		message := gdax.Message{}
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		tickerChan <- message.Price
	}
}

func channelPrinter(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func parseLevel2(message gdax.Message) [][]byte {
	bids, err := json.Marshal(message.Bids)
	if err != nil {
		panic(err)
	}

	asks, err := json.Marshal(message.Asks)
	if err != nil {
		panic(err)
	}

	output := [][]byte{bids, asks}
	return output
}
