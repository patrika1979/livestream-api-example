package main

import (
	"encoding/json"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"log"
	"os"
	"os/signal"
)

type KaikoMessage struct {
	Event   string  `json:"event"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Subscription Subscription
	Data         []Trade
}

type Subscription struct {
	Topic           string `json:"topic"`
	DataVersion     string `json:"data_version"`
	Exchange        string `json:"exchange"`
	InstrumentClass string `json:"instrument_class"`
	Instrument      string `json:"instrument"`
}

type Trade struct {
	Amount    json.Number `json:"amount"`
	Price     json.Number `json:"price"`
	TakerSide bool        `json:"taker_side_sell"`
	Timestamp json.Number `json:"timestamp"`
	TradeId   string      `json:"trade_id"`
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	api_key := getenv("KAIKO_API_KEY", "__REPLACE_ME__")

	exchange := "bnce"         // Binance exchange
	instrument_class := "spot" // Spot market
	instrument := "*-btc"      // all instrument traded against BTC

	url := fmt.Sprintf("wss://us.market-ws.kaiko.io/v2/data/trades_ws.latest/%s:%s:%s", exchange, instrument_class, instrument)
	socket := gowebsocket.New(url)

	socket.ConnectionOptions = gowebsocket.ConnectionOptions{
		UseSSL: false,
	}

	socket.RequestHeader.Set("Sec-WebSocket-Protocol", fmt.Sprintf("api_key, %s", api_key))

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Recieved connect error ", err)
	}

	socket.OnTextMessage = func(data string, socket gowebsocket.Socket) {
		var message KaikoMessage
		json.Unmarshal([]byte(data), &message)
		if message.Event == "update" {
			log.Println(fmt.Sprintf("Received a trade from %s:%s:%s",
				message.Payload.Subscription.Exchange,
				message.Payload.Subscription.InstrumentClass,
				message.Payload.Subscription.Instrument))

			log.Println(message.Payload.Data)
		}
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}
	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}
