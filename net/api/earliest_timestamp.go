package api

import (
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

	response, err := http.Get(requestUri)

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
