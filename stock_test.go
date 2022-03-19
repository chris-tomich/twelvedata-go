package twelvedata

import (
	"testing"

	"github.com/chris-tomich/twelvedata-go/net/api"
)

func BenchmarkGetStocks(b *testing.B) {
	exchangeData := api.NewExchangesRequest().Request()
	exchanges := GetExchangeList(exchangeData)

	stockData := api.NewStocksRequest(exchanges[0].Name).Request()

	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		GetStockList(stockData)
	}
}
