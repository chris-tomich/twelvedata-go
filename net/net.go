package net

const APIBase = "https://api.twelvedata.com"

type TwelveDataRequest interface {
	Request() []byte
}
