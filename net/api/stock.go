package api

import (
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go"
	"github.com/chris-tomich/twelvedata-go/net"
)

const StocksEndpoint = "/stocks"

func NewStocksRequest(e *twelvedata.Exchange) *StockListRequest {
	return &StockListRequest{
		Exchange: e,
	}
}

type StockListRequest struct {
	Exchange *twelvedata.Exchange
}

func (req *StockListRequest) Request() []byte {
	response, err := http.Get(net.APIBase + StocksEndpoint + "?exchange=" + req.Exchange.Name)

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
