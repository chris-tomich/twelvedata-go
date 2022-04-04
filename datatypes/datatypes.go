package datatypes

type Exchange struct {
	Name     string `csv:"name"`
	Code     string `csv:"code"`
	Country  string `csv:"country"`
	Timezone string `csv:"timezone"`
}

type Stock struct {
	Symbol   string `csv:"symbol"`
	Name     string `csv:"name"`
	Currency string `csv:"currency"`
	Exchange string `csv:"exchange"`
	Country  string `csv:"country"`
	Type     string `csv:"type"`
}
