package adapter

import (
	"io"
	"net/http"
)

const baseURL = "https://f1tv.formula1.com"

func ProxyReq(method string, uri string, body io.Reader) (io.ReadCloser, error) {
	url := baseURL + uri

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	/* var resp interface{}
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return "", err
	} */

	return httpResp.Body, nil
}
