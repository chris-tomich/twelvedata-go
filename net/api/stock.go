package api

import (
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/net"
)

const StocksEndpoint = "/stocks"

func NewStocksRequestCSV(exchange string) *StockListRequestCSV {
	return &StockListRequestCSV{
		Exchange: exchange,
	}
}

func NewStocksRequestJSON(exchange string) *StockListRequestJSON {
	return &StockListRequestJSON{
		Exchange: exchange,
	}
}

type StockListRequestCSV struct {
	Exchange string
}

func (req *StockListRequestCSV) Request() []byte {
	response, err := http.Get(net.APIBase + StocksEndpoint + "?exchange=" + req.Exchange + "&format=CSV&delimiter=,")

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

type StockListRequestJSON struct {
	Exchange string
}

func (req *StockListRequestJSON) Request() []byte {
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
