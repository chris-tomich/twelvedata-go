package twelvedata

import (
	"fmt"
	"testing"
)

func TestGetStocks(t *testing.T) {
	exchanges := GetExchangeList(NewExchangesRequest())

	fmt.Printf("%v", exchanges[0])

	stocks := GetStockList(NewStocksRequest(&exchanges[0]))

	fmt.Printf("%v", len(stocks))
}
