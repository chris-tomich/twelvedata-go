package api

import (
	"io"
	"net/http"

	"github.com/chris-tomich/twelvedata-go/net"
)

const EXCHANGES_ENDPOINT = "/exchanges"

type ExchangeListRequest struct{}

func (req *ExchangeListRequest) Request() []byte {
	response, err := http.Get(net.API_BASE + EXCHANGES_ENDPOINT)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return body
}
