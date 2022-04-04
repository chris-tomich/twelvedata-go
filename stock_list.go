package twelvedata

import (
	"fmt"

	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/jszwec/csvutil"
)

type stocksResponse struct {
	Stocks []datatypes.Stock `json:"data"`
}

type StocksRequest struct {
	Symbol   string
	Exchange string
	Country  string
	Type     string
}

func parseStockList(body []byte) ([]datatypes.Stock, error) {
	data := &stocksResponse{
		Stocks: make([]datatypes.Stock, 0, 10),
	}

	err := csvutil.Unmarshal(body, &data.Stocks)

	if err != nil {
		return nil, fmt.Errorf("issue with parsing the stock list: %w", err)
	}

	return data.Stocks, nil
}
