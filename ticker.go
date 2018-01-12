package binance

import (
	"encoding/json"
	"fmt"
)

type Ticks []*Tick

type Tick struct {
	Symbol 			   string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	BidQty             float64 `json:"bidQty,string"`
	AskPrice           float64 `json:"askPrice,string"`
	AskQty             float64 `json:"askQty,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FristID            int     `json:"fristId"`
	LastID             int     `json:"lastId"`
	Count              int     `json:"count"`
}

// Binance API implementation of Ticker endpoint.
//
// Endpoint:  /api/v1/ticker/24hr
// Method: GET
//
// Example: https://api.binance.com/api/v1/ticker/24hr
//
// Sample Response:
//
/*
[
  {
    "symbol": "ETHBTC",
    "priceChange": "0.00085500",
    "priceChangePercent": "0.969",
    "weightedAvgPrice": "0.08812350",
    "prevClosePrice": "0.08827000",
    "lastPrice": "0.08911800",
    "lastQty": "0.02000000",
    "bidPrice": "0.08912300",
    "bidQty": "10.34100000",
    "askPrice": "0.08927000",
    "askQty": "16.48700000",
    "openPrice": "0.08826300",
    "highPrice": "0.09124000",
    "lowPrice": "0.08500000",
    "volume": "297046.59500000",
    "quoteVolume": "26176.78619628",
    "openTime": 1515651669537,
    "closeTime": 1515738069537,
    "firstId": 19426252,
    "lastId": 19938152,
    "count": 511901
  },
  ]
*/

func (client *Client) GetTickers() (Ticks, error) {

	resp, err := client.do("ticker/24hr", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Ticks, 0)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}

func (client *Client) GetTicker(id string) (*Tick, error) {

	resp, err := client.do("ticker/24hr"+"/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Ticks, 1)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res[0], nil
}
