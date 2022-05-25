package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/chris-tomich/twelvedata-go/net"
)

const StocksEndpoint = "/stocks"

type StocksListRequest struct {
	Exchange *datatypes.Exchange
	APIKey   string
}

func NewStocksRequest(apikey string, exchange *datatypes.Exchange) *StocksListRequest {
	return &StocksListRequest{
		Exchange: exchange,
		APIKey:   apikey,
	}
}

func (req *StocksListRequest) Request() ([]byte, error) {
	requestUri := net.APIBase + StocksEndpoint + "?exchange=" + req.Exchange.Name + "&format=CSV&delimiter=,&apikey=" + req.APIKey
	response, err := http.Get(requestUri)

	if err != nil {
		return nil, fmt.Errorf("issue with requesting the stocks list; URI - '%v': %w", requestUri, err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("issue with requesting the stocks list response body; URI - '%v'; response - '%v': %w", requestUri, response, err)
	}

	return body, nil
}
