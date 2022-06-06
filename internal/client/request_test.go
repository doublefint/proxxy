//Package client contains definitions of client request, response structs
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {

	type testcase struct {
		body    io.Reader
		want    Request
		wantErr bool
	}

	tests := map[string]testcase{

		"Parse_validJSON": {
			body: strings.NewReader(`{
					"method": "GET",
					"url": "http://google.com",
					"headers": { "Authentication": "Basic bG9naW46cGFzc3dvcmQ="}
				}`),
			want: Request{
				Method:  "GET",
				URL:     "http://google.com",
				Headers: map[string]string{"Authentication": "Basic bG9naW46cGFzc3dvcmQ="}},
		},
		"Parse_invalidJSON": {
			body:    strings.NewReader(`{method: "GET"}`),
			wantErr: true,
		},
		"Parse_non_HTTP": {
			body:    strings.NewReader(`{"method": "GET", "url": "https://google.com"}`),
			want:    Request{Method: "GET", URL: "https://google.com"},
			wantErr: true,
		},

		"Parse_invalid_URL": {
			body:    strings.NewReader(`{"method": "GET", "url": "\t\n"}`),
			want:    Request{Method: "GET", URL: "\t\n"},
			wantErr: true,
		},
	}

	t.Parallel()

	for name, tc := range tests {
		name, tc := name, tc //parallel
		t.Run(name, func(t *testing.T) {

			t.Parallel()

			got, err := Parse(tc.body)
			if (err != nil) != tc.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, got, "Parse() valid json")

		})
	}
}

func TestRequest_ToHTTPRequestWithContext(t *testing.T) {
	type testcase struct {
		rq      Request
		want    *http.Request
		wantErr bool
	}

	ctx := context.Background()

	tests := map[string]testcase{

		"ok": func() testcase {

			method := "GET"
			url := "http://google.com"
			key, value := "Accept", "application/json"
			headers := map[string]string{key: value}

			req, err := http.NewRequestWithContext(ctx, method, url, nil)
			require.NoError(t, err)

			req.Header.Add(key, value)

			return testcase{
				rq:   Request{Method: method, URL: url, Headers: headers},
				want: req,
			}

		}(),
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			got, err := tc.rq.ToHTTPRequestWithContext(ctx)
			if (err != nil) != tc.wantErr {
				t.Errorf("Request.ToHTTPRequestWithContext() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got, fmt.Sprintf("Request.ToHTTPRequestWithContext() %s", name))

		})
	}
}
