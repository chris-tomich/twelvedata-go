package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chris-tomich/twelvedata-go/net"
)

const TimeSeriesEndpoint = "/time_series"

type InstrumentType string

const (
	Stock                     InstrumentType = "Stock"
	Index                     InstrumentType = "Index"
	ExchangeTradedFunds       InstrumentType = "ETF"
	RealEstateInvestmentTrust InstrumentType = "REIT"
)

type Interval string

const (
	Minute1  Interval = "1min"
	Minute5  Interval = "5min"
	Minute15 Interval = "15min"
	Minute30 Interval = "30min"
	Minute45 Interval = "45min"
	Hourly1  Interval = "1h"
	Hourly2  Interval = "2h"
	Hourly4  Interval = "4h"
	Daily    Interval = "1day"
	Weekly   Interval = "1week"
	Monthly  Interval = "1month"
)

type TimeSeriesRequest struct {
	Symbol        string
	Interval      Interval
	Exchange      string
	Country       string
	Type          InstrumentType
	OutputSize    int
	Format        string
	Delimiter     string
	APIKey        string
	PrePost       string
	DecimalPlaces int
	Order         string
	Timezone      string
	Date          string
	StartDate     *time.Time
	EndDate       *time.Time
	PreviousClose string
}

func NewTimeSeriesRequest(apikey string, symbol string, interval Interval) *TimeSeriesRequest {
	return &TimeSeriesRequest{
		Type:       Stock,
		Symbol:     symbol,
		Interval:   interval,
		APIKey:     apikey,
		OutputSize: 5000,
	}
}

func formatTwelveDataDateTime(date time.Time) string {
	return date.Format(time.DateTime)
}

func (req *TimeSeriesRequest) Request() ([]byte, error) {
	requestUri := fmt.Sprintf("%v%v?format=CSV&delimiter=,&apikey=%v&symbol=%v&interval=%v&outputsize=%v", net.APIBase, TimeSeriesEndpoint, req.APIKey, req.Symbol, string(req.Interval), req.OutputSize)

	if req.Type != Stock {
		requestUri = fmt.Sprintf("%v&type=%v", requestUri, string(req.Type))
	}

	if req.StartDate != nil {
		requestUri = fmt.Sprintf("%v&start_date=%v", requestUri, formatTwelveDataDateTime(*req.StartDate))

		if req.EndDate != nil {
			requestUri = fmt.Sprintf("%v&end_date=%v", requestUri, formatTwelveDataDateTime(*req.EndDate))
		}
	}

	response, err := http.Get(requestUri)

	if err != nil {
		return nil, fmt.Errorf("issue with requesting the stock time series data; URI - '%v': %w", requestUri, err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("issue with reading the stock time series response body; URI - '%v'; response - '%v': %w", requestUri, response, err)
	}

	return body, nil
}
