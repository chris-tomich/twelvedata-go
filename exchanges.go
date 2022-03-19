package twelvedata

import (
	"encoding/json"
	"io"
	"net/http"
)

const EXCHANGES_ENDPOINT = "/exchanges"

type ExchangeType string

const (
	StockExchange ExchangeType = "stock"
	ETFExchange   ExchangeType = "etf"
	IndexExchange ExchangeType = "index"
)

type Exchange struct {
	Name     string
	Code     string
	Country  string
	Timezone string
}

type exchangesResponse struct {
	Exchanges []Exchange `json:"data"`
}

type ExchangesRequest struct {
	Type    ExchangeType
	Name    string
	Code    string
	Country string
}

func NewExchangesRequest() *ExchangesRequest {
	return &ExchangesRequest{
		Type: StockExchange,
	}
}

func GetExchanges(request *ExchangesRequest) []Exchange {
	response, err := http.Get(API_BASE + EXCHANGES_ENDPOINT)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	data := &exchangesResponse{
		Exchanges: make([]Exchange, 0, 10),
	}

	err = json.Unmarshal(body, data)

	if err != nil {
		panic(err)
	}

	return data.Exchanges
}
