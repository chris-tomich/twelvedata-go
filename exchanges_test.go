package twelvedata

import (
	"fmt"
	"testing"
)

func TestGetExchanges(t *testing.T) {
	exchanges := GetExchanges(NewDefaultRequest())

	fmt.Printf("%v", exchanges)
}
