package twelvedata

import (
	"fmt"
	"testing"
)

func TestGetStocks(t *testing.T) {
	exchanges := GetExchanges(NewExchangesRequest())

	fmt.Printf("%v", exchanges[0])

	stocks := GetStocks(NewStocksRequest(&exchanges[0]))

	fmt.Printf("%v", len(stocks))
}
