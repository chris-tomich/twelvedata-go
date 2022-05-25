package twelvedata

import (
	"fmt"

	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/chris-tomich/twelvedata-go/net/api"
)

type TwelveDataClient struct {
	apiKey string
}

func NewTwelveDataClient(apiKey string) *TwelveDataClient {
	if apiKey == "" {
		panic(fmt.Errorf("apiKey can't be empty"))
	}

	return &TwelveDataClient{
		apiKey,
	}
}

func (client *TwelveDataClient) RequestStockExchanges() ([]datatypes.Exchange, error) {
	exchangesData, err := api.NewExchangesRequest(client.apiKey).Request()

	if err != nil {
		return nil, err
	}

	return parseExchangesList(exchangesData)
}

func (client *TwelveDataClient) RequestETFExchanges() ([]datatypes.Exchange, error) {
	req := api.NewExchangesRequest(client.apiKey)
	req.Type = api.ETFExchange

	exchangesData, err := req.Request()

	if err != nil {
		return nil, err
	}

	return parseExchangesList(exchangesData)
}

func (client *TwelveDataClient) RequestIndexExchanges() ([]datatypes.Exchange, error) {
	req := api.NewExchangesRequest(client.apiKey)
	req.Type = api.IndexExchange

	exchangesData, err := req.Request()

	if err != nil {
		return nil, err
	}

	return parseExchangesList(exchangesData)
}

func (client *TwelveDataClient) RequestStocks(exchange *datatypes.Exchange) ([]datatypes.Stock, error) {
	stocksData, err := api.NewStocksRequest(client.apiKey, exchange).Request()

	if err != nil {
		return nil, err
	}

	return parseStocksList(stocksData)
}

func (client *TwelveDataClient) RequestTimeSeriesData(symbol string, interval api.Interval) ([]datatypes.TimeSeriesData, error) {
	timeSeriesData, err := api.NewTimeSeriesRequest(client.apiKey, symbol, interval).Request()

	if err != nil {
		return nil, err
	}

	return parseTimeSeriesData(timeSeriesData)
}

func (client *TwelveDataClient) RequestEarliestTimestamp(symbol string, interval api.Interval) (*datatypes.EarliestTimestamp, error) {
	earliestTimestampData, err := api.NewEarliestTimestampRequest(client.apiKey, symbol, interval).Request()

	if err != nil {
		return nil, err
	}

	return parseEarliestTimestamp(earliestTimestampData)
}
