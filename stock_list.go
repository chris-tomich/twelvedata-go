package twelvedata

import "github.com/jszwec/csvutil"

type Stock struct {
	Symbol   string `csv:"symbol"`
	Name     string `csv:"name"`
	Currency string `csv:"currency"`
	Exchange string `csv:"exchange"`
	Country  string `csv:"country"`
	Type     string `csv:"type"`
}

type stocksResponse struct {
	Stocks []Stock `json:"data"`
}

type StocksRequest struct {
	Symbol   string
	Exchange string
	Country  string
	Type     string
}

func GetStockList(body []byte) []Stock {
	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err := csvutil.Unmarshal(body, &data.Stocks)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}
