package binance

import (
	"encoding/json"
	"fmt"
)

type Ticks []*Tick

type Tick struct {
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	AskPrice           float64 `json:"askPrice,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	OpenTime           int64  `json:"openTime,string"`
	CloseTime          int64  `json:"closeTime,string"`
	FristID            int    `json:"fristId,string"`
	LastID             int    `json:"lastId,string"`
	Count              int    `json:"count,string"`
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
		  "priceChange": "-94.99999800",
		  "priceChangePercent": "-95.960",
		  "weightedAvgPrice": "0.29628482",
		  "prevClosePrice": "0.10002000",
		  "lastPrice": "4.00000200",
		  "bidPrice": "4.00000000",
		  "askPrice": "4.00000200",
		  "openPrice": "99.00000000",
		  "highPrice": "100.00000000",
		  "lowPrice": "0.10000000",
		  "volume": "8913.30000000",
		  "openTime": 1499783499040,
		  "closeTime": 1499869899040,
		  "fristId": 28385,   // First tradeId
		  "lastId": 28460,    // Last tradeId
		  "count": 76         // Trade count
		}
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
