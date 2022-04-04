package twelvedata

import (
	"fmt"

	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/jszwec/csvutil"
)

type exchangesResponse struct {
	Exchanges []datatypes.Exchange `json:"data"`
}

func parseExchangeList(body []byte) ([]datatypes.Exchange, error) {
	data := &exchangesResponse{
		Exchanges: make([]datatypes.Exchange, 0, 10),
	}

	err := csvutil.Unmarshal(body, &data.Exchanges)

	if err != nil {
		return nil, fmt.Errorf("issue with parsing the exchange list: %w", err)
	}

	return data.Exchanges, nil
}
