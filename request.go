package waxpeer

import (
	"github.com/valyala/fasthttp"
)

func Get(url string) (*[]byte, error) {
	request := &fasthttp.Request{}
	request.Header.SetRequestURI(url)
	request.Header.SetMethod("GET")
	response := &fasthttp.Response{}
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	b := response.Body()
	return &b, nil
}
