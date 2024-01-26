package api

import (
	"crypto/tls"
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
	MICode  string
	Country string
	APIKey  string
}

func NewExchangesRequest(apikey string) *ExchangesListRequest {
	return &ExchangesListRequest{
		Type:   StockExchange,
		APIKey: apikey,
	}
}

func (req *ExchangesListRequest) Request() ([]byte, error) {
	requestUri := net.APIBase + ExchangesEndpoint + "?format=CSV&delimiter=,&apikey=" + req.APIKey

	if req.Type != StockExchange {
		requestUri += "&type=" + string(req.Type)
	}

	// TODO: Remove this when TwelveData fixes the TLS for their API
	// Create a custom HTTP client that skips TLS verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Restore this call when TwelveData fixes the TLS for their API
	// response, err := http.Get(requestUri)
	response, err := client.Get(requestUri)

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
