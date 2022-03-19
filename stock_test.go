package twelvedata

import (
	"fmt"
	"testing"

	"github.com/chris-tomich/twelvedata-go/net/api"
)

func TestGetStocks(t *testing.T) {
	exchanges := GetExchangeList(api.NewExchangesRequest())

	fmt.Printf("%v", exchanges[0])

	stocks := GetStockList(api.NewStocksRequest(exchanges[0].Name))

	fmt.Printf("%v", len(stocks))
}
