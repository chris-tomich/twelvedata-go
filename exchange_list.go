package twelvedata

import (
	"encoding/json"
)

type Exchange struct {
	Name     string
	Code     string
	Country  string
	Timezone string
}

type exchangesResponse struct {
	Exchanges []Exchange `json:"data"`
}

func GetExchangeList(body []byte) []Exchange {
	data := &exchangesResponse{
		Exchanges: make([]Exchange, 0, 10),
	}

	err := json.Unmarshal(body, data)

	if err != nil {
		panic(err)
	}

	return data.Exchanges
}
