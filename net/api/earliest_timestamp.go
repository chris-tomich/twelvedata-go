package api

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/net"
)

const EarliestTimestampEndpoint = "/earliest_timestamp"

type EarliestTimestampRequest struct {
	Symbol   string
	Interval Interval
	APIKey   string
}

func NewEarliestTimestampRequest(apikey string, symbol string, interval Interval) *EarliestTimestampRequest {
	return &EarliestTimestampRequest{
		Symbol:   symbol,
		Interval: interval,
		APIKey:   apikey,
	}
}

func (req *EarliestTimestampRequest) Request() ([]byte, error) {
	requestUri := net.APIBase + EarliestTimestampEndpoint + "?apikey=" + req.APIKey + "&symbol=" + req.Symbol + "&interval=" + string(req.Interval)

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
		return nil, fmt.Errorf("issue with requesting the earliest timestamp data; URI - '%v': %w", requestUri, err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("issue with requesting the earliest timestamp data; URI - '%v'; response - '%v': %w", requestUri, response, err)
	}

	return body, nil
}
