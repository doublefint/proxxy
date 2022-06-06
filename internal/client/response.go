package client

import (
	"net/http"
)

//Response is client response struct
type Response struct {
	ID      int64             `json:"id"`      // generated unique id
	Length  int64             `json:"length"`  // content length of 3rd-party service response
	Status  string            `json:"status"`  // HTTP status of 3rd-party service response
	Headers map[string]string `json:"headers"` // headers array from 3rd-party service response
}

//NewResponse creates client response from http response
func NewResponse(id int64, resp *http.Response) Response {
	cr := Response{ID: id, Status: resp.Status, Length: resp.ContentLength, Headers: map[string]string{}}
	for k, v := range resp.Header {
		for _, header := range v {
			cr.Headers[k] = header
		}
	}
	return cr
}
