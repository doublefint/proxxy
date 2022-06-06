//Package client contains definitions of client request, response structs
package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	//ErrNonHTTP raised for non http url
	ErrNonHTTP = errors.New("unexpected protocol")
)

//Request is client request struct
type Request struct {
	Method  string            `json:"method" example:"GET"`            //client request method
	URL     string            `json:"url" example:"http://google.com"` //client request url
	Headers map[string]string `json:"headers"`                         //client request headers
}

//Parse parses the contents of the request and returns client.Request
func Parse(body io.Reader) (Request, error) {

	var data Request

	bytes, err := io.ReadAll(body)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return data, err
	}

	u, err := url.Parse(data.URL)
	if err != nil {
		return data, err
	}

	if u.Scheme != "http" {
		return data, ErrNonHTTP
	}

	return data, nil
}

//ToHTTPRequestWithContext creates new http request
func (rq Request) ToHTTPRequestWithContext(ctx context.Context) (*http.Request, error) {

	req, err := http.NewRequestWithContext(ctx, rq.Method, rq.URL, nil)
	if err != nil {
		return req, err
	}
	for key, value := range rq.Headers {
		req.Header.Add(key, value)
	}
	return req, err

}
