package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testcase struct {
		id   int64
		hr   *http.Response
		want Response
	}

	tests := map[string]testcase{
		"main": func() testcase {

			id := int64(1)
			status := "200 OK"
			hr := &http.Response{Status: status, Header: http.Header{}}
			key, value := "Content-Type", "application/json"
			hr.Header.Add(key, value)

			return testcase{
				id: id,
				hr: hr,
				want: Response{
					ID:      id,
					Status:  status,
					Headers: map[string]string{key: value},
				},
			}

		}(),
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewResponse(tc.id, tc.hr)
			assert.Equal(t, tc.want, got, "http.Response to client response")
		})
	}
}
