package twelvedata

import (
	"encoding/json"
	"fmt"

	"github.com/chris-tomich/twelvedata-go/datatypes"
)

func parseEarliestTimestamp(body []byte) (*datatypes.EarliestTimestamp, error) {
	earliestTimestamp := &datatypes.EarliestTimestamp{}

	err := json.Unmarshal(body, earliestTimestamp)

	if err != nil {
		return nil, fmt.Errorf("issue with parsing the earliest timestamp data: %w;\nJSON body: %v", err, string(body))
	}

	return earliestTimestamp, nil
}
