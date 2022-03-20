package twelvedata

import (
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

func GetExchangeList(body []byte) []Exchange {
	data := &exchangesResponse{
		Exchanges: make([]Exchange, 0, 10),
	}

	err := csvutil.Unmarshal(body, &data.Exchanges)

	if err != nil {
		panic(err)
	}

	return data.Exchanges
}
