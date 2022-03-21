package twelvedata

import (
	"testing"

	"github.com/chris-tomich/twelvedata-go/net/api"
)

func BenchmarkGetStocksCSV(b *testing.B) {
	exchangeData := api.NewExchangesRequest().Request()
	exchanges := GetExchangeList(exchangeData)

	stockData := api.NewStocksRequestCSV(exchanges[0].Name).Request()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		GetStockListCSV(stockData)
	}
}

func BenchmarkGetStocksJSON(b *testing.B) {
	exchangeData := api.NewExchangesRequest().Request()
	exchanges := GetExchangeList(exchangeData)

	stockData := api.NewStocksRequestJSON(exchanges[0].Name).Request()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		GetStockListJSON(stockData)
	}
}

func BenchmarkGetStocksDynamicCSV(b *testing.B) {
	exchangeData := api.NewExchangesRequest().Request()
	exchanges := GetExchangeList(exchangeData)

	stockData := api.NewStocksRequestCSV(exchanges[0].Name).Request()

	b.ReportAllocs()

	u := &CSVUnmarshaller{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		GetStockListDynamicDispatch(u, stockData)
	}
}

func BenchmarkGetStocksStaticCSV(b *testing.B) {
	exchangeData := api.NewExchangesRequest().Request()
	exchanges := GetExchangeList(exchangeData)

	stockData := api.NewStocksRequestCSV(exchanges[0].Name).Request()

	b.ReportAllocs()

	u := &CSVUnmarshaller{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		GetStockListStaticDispatch(u, stockData)
	}
}
