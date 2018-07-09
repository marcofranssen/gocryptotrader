package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeInfoResponse struct {
	Symbols []Symbol `json:"symbols"`
}

type Symbol struct {
	Name       string `json:"symbol"`
	Status     string `json:"status"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func createConfigString(buffer bytes.Buffer) {

}

func main() {
	resp, err := http.Get("http://api.binance.com/api/v1/exchangeInfo")
	panicOnError(err)
	defer resp.Body.Close()

	var data ExchangeInfoResponse
	json.NewDecoder(resp.Body).Decode(&data)

	var buffer bytes.Buffer
	for index, symbol := range data.Symbols {
		if symbol.Status == "TRADING" {
			buffer.WriteString(symbol.BaseAsset + "-" + symbol.QuoteAsset)
			if index < len(data.Symbols)-1 {
				buffer.WriteString(",")
			}
		}
	}

	fmt.Println()
	fmt.Println("Copy following Binance trading pairs to your config.json:")
	fmt.Println()
	fmt.Print(buffer.String())
}
