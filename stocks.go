package twelvedata

import (
	"encoding/json"
	"io"
	"net/http"
)

const STOCKS_ENDPOINT = "/stocks"

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

func NewStocksRequest(e *Exchange) *StocksRequest {
	return &StocksRequest{
		Exchange: e.Name,
	}
}

func GetStocks(request *StocksRequest) []Stock {
	response, err := http.Get(API_BASE + STOCKS_ENDPOINT + "?exchange=" + request.Exchange)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err = json.Unmarshal(body, data)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}
