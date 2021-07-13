package waxpeer

import (
	"github.com/valyala/fasthttp"
)

func Get(url string) (*[]byte, error) {
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(url)
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	b := response.Body()
	return &b, nil
}

func Post(url string, body []byte) (*[]byte, error) {
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(url)
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	request.SetBody(body)
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	b := response.Body()
	return &b, nil
}
