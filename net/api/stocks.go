package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/net"
)

const StocksEndpoint = "/stocks"

type StocksListRequest struct {
	exchangeCode string
	apiKey       string
}

func NewStocksRequest(apikey string, exchangeCode string) *StocksListRequest {
	return &StocksListRequest{
		exchangeCode: exchangeCode,
		apiKey:       apikey,
	}
}

func (req *StocksListRequest) Request() ([]byte, error) {
	requestUri := net.APIBase + StocksEndpoint + "?mic_code=" + req.exchangeCode + "&format=CSV&delimiter=,&apikey=" + req.apiKey
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
