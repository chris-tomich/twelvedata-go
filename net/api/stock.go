package api

import (
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/net"
)

const StocksEndpoint = "/stocks"

func NewStocksRequest(exchange string) *StockListRequest {
	return &StockListRequest{
		Exchange: exchange,
	}
}

type StockListRequest struct {
	Exchange string
}

func (req *StockListRequest) Request() []byte {
	response, err := http.Get(net.APIBase + StocksEndpoint + "?exchange=" + req.Exchange)

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
