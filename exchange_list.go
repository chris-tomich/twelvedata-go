package twelvedata

import (
	"fmt"

	"github.com/jszwec/csvutil"
)

type Exchange struct {
	Name     string `csv:"name"`
	Code     string `csv:"code"`
	Country  string `csv:"country"`
	Timezone string `csv:"timezone"`
}

type exchangesResponse struct {
	Exchanges []Exchange `json:"data"`
}

func parseExchangeList(body []byte) ([]Exchange, error) {
	data := &exchangesResponse{
		Exchanges: make([]Exchange, 0, 10),
	}

	err := csvutil.Unmarshal(body, &data.Exchanges)

	if err != nil {
		return nil, fmt.Errorf("issue with parsing the exchange list: %w", err)
	}

	return data.Exchanges, nil
}
