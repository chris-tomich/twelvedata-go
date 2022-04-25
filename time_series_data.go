package twelvedata

import (
	"fmt"

	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/jszwec/csvutil"
)

func parseTimeSeriesData(body []byte) ([]datatypes.TimeSeriesData, error) {
	data := make([]datatypes.TimeSeriesData, 0, 30)

	err := csvutil.Unmarshal(body, &data)

	if err != nil {
		return nil, fmt.Errorf("issue with parsing the time series data: %w", err)
	}

	return data, nil
}
