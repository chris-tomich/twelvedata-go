package twelvedata

import (
	"encoding/json"

	"github.com/chris-tomich/twelvedata-go/net"
)

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

func GetExchangeList(request net.TwelveDataRequest) []Exchange {
	body := request.Request()

	data := &exchangesResponse{
		Exchanges: make([]Exchange, 0, 10),
	}

	err := json.Unmarshal(body, data)

	if err != nil {
		panic(err)
	}

	return data.Exchanges
}
