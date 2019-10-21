package btcmarkets

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
)

// Response is the genralized response type
type Response struct {
	Success         bool              `json:"success"`
	ErrorCode       int               `json:"errorCode"`
	ErrorMessage    string            `json:"errorMessage"`
	ID              int               `json:"id"`
	Responses       []ResponseDetails `json:"responses"`
	ClientRequestID string            `json:"clientRequestId"`
	Orders          []Order           `json:"orders"`
	Status          string            `json:"status"`
}

// ResponseDetails holds order status details
type ResponseDetails struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ID           int64  `json:"id"`
}

// Market holds a tradable market instrument
type Market struct {
	Instrument string `json:"instrument"`
	Currency   string `json:"currency"`
}

// Ticker holds ticker information
type Ticker struct {
	BestAsk    float64       `json:"bestAsk"`
	BestBid    float64       `json:"bestBid"`
	Currency   currency.Code `json:"currency"`
	High24h    float64       `json:"high24h"`
	Instrument currency.Pair `json:"instrument"`
	LastPrice  float64       `json:"lastPrice"`
	Low24h     float64       `json:"low24h"`
	Price24h   float64       `json:"price24h"`
	Timestamp  int64         `json:"timestamp"`
	Volume24h  float64       `json:"volume24h"`
}

// Orderbook holds current orderbook information returned from the exchange
type Orderbook struct {
	Currency   string      `json:"currency"`
	Instrument string      `json:"instrument"`
	Timestamp  int64       `json:"timestamp"`
	Asks       [][]float64 `json:"asks"`
	Bids       [][]float64 `json:"bids"`
}

// Trade holds trade information
type Trade struct {
	TradeID int64   `json:"tid"`
	Amount  float64 `json:"amount"`
	Price   float64 `json:"price"`
	Date    int64   `json:"date"`
}

// TradingFee 30 day trade volume
type TradingFee struct {
	Success        bool    `json:"success"`
	ErrorCode      int     `json:"errorCode"`
	ErrorMessage   string  `json:"errorMessage"`
	TradingFeeRate float64 `json:"tradingfeerate"`
	Volume30Day    float64 `json:"volume30day"`
}

// OrderToGo holds order information to be sent to the exchange
type OrderToGo struct {
	Currency        string `json:"currency"`
	Instrument      string `json:"instrument"`
	Price           int64  `json:"price"`
	Volume          int64  `json:"volume"`
	OrderSide       string `json:"orderSide"`
	OrderType       string `json:"ordertype"`
	ClientRequestID string `json:"clientRequestId"`
}

// Order holds order information
type Order struct {
	ID              int64           `json:"id"`
	Currency        string          `json:"currency"`
	Instrument      string          `json:"instrument"`
	OrderSide       string          `json:"orderSide"`
	OrderType       string          `json:"ordertype"`
	CreationTime    float64         `json:"creationTime"`
	Status          string          `json:"status"`
	ErrorMessage    string          `json:"errorMessage"`
	Price           float64         `json:"price"`
	Volume          float64         `json:"volume"`
	OpenVolume      float64         `json:"openVolume"`
	ClientRequestID string          `json:"clientRequestId"`
	Trades          []TradeResponse `json:"trades"`
}

// TradeResponse holds trade information
type TradeResponse struct {
	ID           int64   `json:"id"`
	CreationTime float64 `json:"creationTime"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Volume       float64 `json:"volume"`
	Fee          float64 `json:"fee"`
}

// AccountBalance holds account balance details
type AccountBalance struct {
	Balance      float64 `json:"balance"`
	PendingFunds float64 `json:"pendingFunds"`
	Currency     string  `json:"currency"`
}

// WithdrawRequestCrypto is a generalized withdraw request type
type WithdrawRequestCrypto struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Address  string `json:"address"`
}

// WithdrawRequestAUD is a generalized withdraw request type
type WithdrawRequestAUD struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	BSBNumber     string `json:"bsbNumber"`
}

// WithdrawalFees the large list of predefined withdrawal fees
// Prone to change
var WithdrawalFees = map[currency.Code]float64{
	currency.AUD:  0,
	currency.BTC:  0.001,
	currency.ETH:  0.001,
	currency.ETC:  0.001,
	currency.LTC:  0.0001,
	currency.XRP:  0.15,
	currency.BCH:  0.0001,
	currency.OMG:  0.15,
	currency.POWR: 5,
}

// WsSubscribe message sent via ws to subscribe
type WsSubscribe struct {
	MarketIDs   []string `json:"marketIds,omitempty"`
	Channels    []string `json:"channels"`
	MessageType string   `json:"messageType"`
}

// WsMessageType message sent via ws to determine type
type WsMessageType struct {
	MessageType string `json:"messageType"`
}

// WsTick message received for ticker data
type WsTick struct {
	Currency    string    `json:"marketId"`
	Timestamp   time.Time `json:"timestamp"`
	Bid         float64   `json:"bestBid,string"`
	Ask         float64   `json:"bestAsk,string"`
	Last        float64   `json:"lastPrice,string"`
	Volume      float64   `json:"volume24h,string"`
	Price24h    float64   `json:"price24h,string"`
	Low24h      float64   `json:"low24h,string"`
	High24      float64   `json:"high24h,string"`
	MessageType string    `json:"messageType"`
}

// WsTrade message received for trade data
type WsTrade struct {
	Currency    string    `json:"marketId"`
	Timestamp   time.Time `json:"timestamp"`
	TradeID     int64     `json:"tradeId"`
	Price       float64   `json:"price,string"`
	Volume      float64   `json:"volume,string"`
	MessageType string    `json:"messageType"`
}

// WsOrderbook message received for orderbook data
type WsOrderbook struct {
	Currency    string     `json:"marketId"`
	Timestamp   time.Time  `json:"timestamp"`
	Bids        [][]string `json:"bids"`
	Asks        [][]string `json:"asks"`
	MessageType string     `json:"messageType"`
}

type WsError struct {
	MessageType string `json:"messageType"`
	Code        int64  `json:"code"`
	Message     string `json:"message"`
}
