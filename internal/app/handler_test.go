//go:build integration

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"proxxy/internal/client"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApp_rootHandler(t *testing.T) {

	logger := log.New(os.Stdout, "test> ", log.LstdFlags)
	srv := New("8080", logger)
	ctx := context.Background()
	handler := srv.rootHandler(ctx)

	w := httptest.NewRecorder()
	body := strings.NewReader(`{
		"method": "GET",
		"url": "http://google.com",
		"headers": { "Accept": "*/*" }
	}`)

	r := httptest.NewRequest(http.MethodGet, "/", body)
	handler(w, r)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode, "OK")
	assert.True(t, srv.counter > 0, "counter incremented")

	bytes, err := io.ReadAll(w.Result().Body)
	require.NoError(t, err, "response has data")

	var clresp client.Response
	err = json.Unmarshal(bytes, &clresp)
	require.NoError(t, err, "response parsed")

	want := fmt.Sprintf("%d %s", http.StatusMovedPermanently, http.StatusText(http.StatusMovedPermanently))
	assert.Equal(t, want, clresp.Status, "google moved permanently to https")

}
