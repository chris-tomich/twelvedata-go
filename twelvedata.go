package twelvedata

import "github.com/chris-tomich/twelvedata-go/net/api"

type TwelveDataClient struct {
	apiKey string
}

func NewTwelveDataClient(apiKey string) *TwelveDataClient {
	return &TwelveDataClient{
		apiKey,
	}
}

func (client *TwelveDataClient) Exchanges() ([]Exchange, error) {
	exchangesData, err := api.NewExchangesRequest().Request()

	if err != nil {
		return nil, err
	}

	return parseExchangeList(exchangesData)
}

func (client *TwelveDataClient) Stocks(exchange string) ([]Stock, error) {
	stocksData, err := api.NewStocksRequest(exchange).Request()

	if err != nil {
		return nil, err
	}

	return parseStockList(stocksData)
}
