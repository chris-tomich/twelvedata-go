package twelvedata

import (
	"encoding/json"

	"github.com/chris-tomich/twelvedata-go/net"
)

type Stock struct {
	Symbol   string
	Name     string
	Currency string
	Exchange string
	MicCode  string `json:"mic_code"`
	Country  string
	Type     string
}

type stocksResponse struct {
	Stocks []Stock `json:"data"`
}

type StocksRequest struct {
	Symbol   string
	Exchange string
	Country  string
	Type     string
}

func GetStockList(request net.TwelveDataRequest) []Stock {
	body := request.Request()

	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err := json.Unmarshal(body, data)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}
