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
	exchangeData, err := api.NewExchangesRequest().Request()

	if err != nil {
		return nil, err
	}

	return parseExchangeList(exchangeData)
}
