package app

import (
	"net"
	"net/http"
	"time"
)

const (
	clientTimeout        = 10
	dialTimeout          = 5
	dialKeepAliveTimeout = 15
	// ct - clientTransport
	ctTLSHandshakeTimeout   = 5
	ctResponseHeaderTimeout = 5
	ctExpectContinueTimeout = 5
	ctMaxConnsPerHost       = 100
	ctMaxIdleConns          = 100
	ctMaxIdleConnsPerHost   = 100
)

func httpClient() *http.Client {

	return &http.Client{
		Timeout: clientTimeout * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   dialTimeout * time.Second,
				KeepAlive: dialKeepAliveTimeout * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   ctTLSHandshakeTimeout * time.Second,
			ResponseHeaderTimeout: ctResponseHeaderTimeout * time.Second,
			ExpectContinueTimeout: ctExpectContinueTimeout * time.Second,
			MaxIdleConns:          int(ctMaxIdleConns),
			MaxConnsPerHost:       int(ctMaxConnsPerHost),
			MaxIdleConnsPerHost:   int(ctMaxIdleConnsPerHost),
			//NOTE: DisableCompression:    true,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

}
