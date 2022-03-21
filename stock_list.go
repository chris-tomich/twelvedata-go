package twelvedata

import (
	"encoding/json"

	"github.com/jszwec/csvutil"
)

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

func GetStockListJSON(body []byte) []Stock {
	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err := json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}

func GetStockListCSV(body []byte) []Stock {
	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err := csvutil.Unmarshal(body, &data.Stocks)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}

type Unmarshaller interface {
	Unmarshal(data []byte, v interface{}) error
}

type CSVUnmarshaller struct{}

func (*CSVUnmarshaller) Unmarshal(data []byte, v interface{}) error {
	return csvutil.Unmarshal(data, v)
}

type JSONUnmarshaller struct{}

func (*JSONUnmarshaller) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func GetStockListDynamicDispatch(unmarshaller Unmarshaller, body []byte) []Stock {
	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err := unmarshaller.Unmarshal(body, &data.Stocks)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}

func GetStockListStaticDispatch[U Unmarshaller](unmarshaller U, body []byte) []Stock {
	data := &stocksResponse{
		Stocks: make([]Stock, 0, 10),
	}

	err := unmarshaller.Unmarshal(body, &data.Stocks)

	if err != nil {
		panic(err)
	}

	return data.Stocks
}
