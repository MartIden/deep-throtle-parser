package request

import "net/http"

func GetPageByUri(uri string) (*http.Response, error) {
	client := http.Client{}
	resp, err := client.Get(uri)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
