package datatypes

type Exchange struct {
	Name     string `csv:"name"`
	MICode   string `csv:"code"`
	Country  string `csv:"country"`
	Timezone string `csv:"timezone"`
}

type Stock struct {
	Symbol   string `csv:"symbol"`
	Name     string `csv:"name"`
	Currency string `csv:"currency"`
	Exchange string `csv:"exchange"`
	MICode   string `csv:"mic_code"`
	Country  string `csv:"country"`
	Type     string `csv:"type"`
}

type TimeSeriesData struct {
	DateTime string `csv:"datetime"`
	Open     string `csv:"open"`
	High     string `csv:"high"`
	Low      string `csv:"low"`
	Close    string `csv:"close"`
	Volume   string `csv:"volume,omitempty"`
}

type EarliestTimestamp struct {
	DateTime string `json:"datetime"`
	UnixTime int64  `json:"unix_time"`
}
