package twelvedata

import (
	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/chris-tomich/twelvedata-go/net/api"
)

type TwelveDataClient struct {
	apiKey string
}

func NewTwelveDataClient(apiKey string) *TwelveDataClient {
	return &TwelveDataClient{
		apiKey,
	}
}

func (client *TwelveDataClient) RequestStockExchanges() ([]datatypes.Exchange, error) {
	exchangesData, err := api.NewExchangesRequest().Request()

	if err != nil {
		return nil, err
	}

	return parseExchangesList(exchangesData)
}

func (client *TwelveDataClient) RequestETFExchanges() ([]datatypes.Exchange, error) {
	req := api.NewExchangesRequest()
	req.Type = api.ETFExchange

	exchangesData, err := req.Request()

	if err != nil {
		return nil, err
	}

	return parseExchangesList(exchangesData)
}

func (client *TwelveDataClient) RequestIndexExchanges() ([]datatypes.Exchange, error) {
	req := api.NewExchangesRequest()
	req.Type = api.IndexExchange

	exchangesData, err := req.Request()

	if err != nil {
		return nil, err
	}

	return parseExchangesList(exchangesData)
}

func (client *TwelveDataClient) RequestStocks(exchange *datatypes.Exchange) ([]datatypes.Stock, error) {
	stocksData, err := api.NewStocksRequest(exchange).Request()

	if err != nil {
		return nil, err
	}

	return parseStocksList(stocksData)
}
