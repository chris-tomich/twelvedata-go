package net

const API_BASE = "https://api.twelvedata.com"

type TwelveDataRequest interface {
	Request() []byte
}
