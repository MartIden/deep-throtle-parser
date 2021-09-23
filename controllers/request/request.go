package request

import (
	"net/http"
	"net/url"
	"time"
)

type HTTPClientSettings struct {
	httpType string
	uri      string
	headers  map[string]string
	proxy    url.URL
	cookies  http.CookieJar
	timeout  time.Duration
}

func GetPageByUri(uri string) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Get(uri)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
