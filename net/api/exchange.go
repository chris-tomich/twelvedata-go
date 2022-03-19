package api

import (
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/net"
)

const ExchangesEndpoint = "/exchanges"

type ExchangeType string

const (
	StockExchange ExchangeType = "stock"
	ETFExchange   ExchangeType = "etf"
	IndexExchange ExchangeType = "index"
)

func NewExchangesRequest() *ExchangeListRequest {
	return &ExchangeListRequest{
		Type: StockExchange,
	}
}

type ExchangeListRequest struct {
	Type    ExchangeType
	Name    string
	Code    string
	Country string
}

func (req *ExchangeListRequest) Request() []byte {
	response, err := http.Get(net.APIBase + ExchangesEndpoint)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return body
}
