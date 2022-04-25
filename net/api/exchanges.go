package api

import (
	"fmt"
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

type ExchangesListRequest struct {
	Type    ExchangeType
	Name    string
	Code    string
	Country string
}

func NewExchangesRequest() *ExchangesListRequest {
	return &ExchangesListRequest{
		Type: StockExchange,
	}
}

func (req *ExchangesListRequest) Request() ([]byte, error) {
	requestUri := net.APIBase + ExchangesEndpoint + "?format=CSV&delimiter=,"

	if req.Type != StockExchange {
		requestUri += "&type=" + string(req.Type)
	}

	response, err := http.Get(requestUri)

	if err != nil {
		return nil, fmt.Errorf("issue with requesting the exchanges list; URI - '%v': %w", requestUri, err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("issue with reading the exchanges list response body; URI - '%v'; response - '%v': %w", requestUri, response, err)
	}

	return body, nil
}
