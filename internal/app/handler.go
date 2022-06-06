package app

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"proxxy/internal/client"
)

// Proxxy endpoint
// @Summary Proxxy endpoint
// @Description Proxying **HTTP**-requests to 3rd-party services.
// @Tags /
// @Accept json
// @Produce json
// @Param  request  body  client.Request  true  "client request"
// @Success 200 {object} client.Response
// @Failure 400 {string} string "invalid input"
// @Failure 500 {string} string "server side error"
// @Router / [get]
func (a *App) rootHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := a.inc()

		a.logger.Printf("[%d] %s /\n", id, r.Method)

		crq, err := client.Parse(r.Body) // client request
		if err != nil {
			a.badRequest(w, id, "failed to parse client request", err)
			return
		}

		req, err := crq.ToHTTPRequestWithContext(ctx)
		if err != nil {
			a.badRequest(w, id, "failed to convert to http request", err)
			return
		}

		resp, err := httpClient().Do(req)
		if resp != nil && resp.Body != nil { //NOTE: redirect

			defer func() {
				_, err := io.Copy(io.Discard, resp.Body)
				if err != nil {
					a.logger.Printf("failed to read resp.Body %v", err)
				}
				if err := resp.Body.Close(); err != nil {
					a.logger.Printf("failed to close resp.Body %v", err)
				}
			}()
		}

		if err != nil {
			a.logger.Printf("[%d] failed to send request %s %s: %v\n", id, crq.Method, crq.URL, err)
			http.Error(w, "failed to send request", http.StatusInternalServerError)
			return
		}

		rsp := client.NewResponse(id, resp)
		a.m.store(id, mapItem{rq: crq, rsp: rsp})

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			http.Error(w, "failed to output json", http.StatusInternalServerError)
		}

	}
}

func (a *App) badRequest(w http.ResponseWriter, id int64, msg string, err error) {
	a.logger.Printf("[%d] %s %v\n", id, msg, err)
	http.Error(w, msg, http.StatusBadRequest)
}
