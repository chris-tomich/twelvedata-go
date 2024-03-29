package twelvedata

import (
	"testing"

	"github.com/chris-tomich/twelvedata-go/net/api"
)

func BenchmarkParseStocks(b *testing.B) {
	exchangeData, err := api.NewExchangesRequest("").Request()

	if err != nil {
		panic(err)
	}

	exchanges, err := parseExchangesList(exchangeData)

	if err != nil {
		panic(err)
	}

	stockData, err := api.NewStocksRequest("", exchanges[0].Name).Request()

	if err != nil {
		panic(err)
	}

	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		parseStocksList(stockData)
	}
}
