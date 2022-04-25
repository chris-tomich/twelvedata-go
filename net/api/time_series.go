package api

const TimeSeriesEndpoint = "/time_series"

type InstrumentType string

const (
	Stock                     InstrumentType = "Stock"
	Index                     InstrumentType = "Index"
	ExchangeTradedFunds       InstrumentType = "ETF"
	RealEstateInvestmentTrust InstrumentType = "REIT"
)

type TimeSeriesRequest struct {
	Symbol        string
	Interval      string
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
	StartDate     string
	EndDate       string
	PreviousClose string
}

func NewTimeSeriesRequest() *TimeSeriesRequest {
	return &TimeSeriesRequest{}
}
