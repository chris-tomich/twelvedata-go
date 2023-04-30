package twelvedata

import (
	"encoding/json"
	"fmt"

	"github.com/chris-tomich/twelvedata-go/datatypes"
	"github.com/jszwec/csvutil"
)

func parseTimeSeriesData(body []byte) ([]datatypes.TimeSeriesData, error) {
	data := make([]datatypes.TimeSeriesData, 0, 30)

	err := csvutil.Unmarshal(body, &data)

	if err != nil {
		unmarshalledErrorJson := make(map[string]any)
		unmarshalledError := json.Unmarshal(body, &unmarshalledErrorJson)

		if unmarshalledError != nil {
			return nil, fmt.Errorf("issue with parsing the time series data: %w;\nbody: %v", err, string(body))
		}

		code, ok := unmarshalledErrorJson["code"].(float64)

		if !ok {
			return nil, fmt.Errorf("issue with unmarshalling code for the time series body: %v, unmarshalled JSON: %v", string(body), unmarshalledErrorJson)
		}

		message, ok := unmarshalledErrorJson["message"].(string)

		if !ok {
			return nil, fmt.Errorf("issue with unmarshalling message for the time series body: %v, unmarshalled JSON: %v", string(body), unmarshalledErrorJson)
		}

		if int64(code) == 400 && message == "No data is available on the specified dates. Try setting different start/end dates." {
			return make([]datatypes.TimeSeriesData, 0), nil
		} else {
			return nil, fmt.Errorf("unexpected error when parsing the time series body: %v", string(body))
		}
	}

	return data, nil
}
