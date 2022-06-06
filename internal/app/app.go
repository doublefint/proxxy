//Package app with proxxy logic
package app

import (
	"log"
	"net/http"
	"sync/atomic"
)

//App is proxxy server
type App struct {
	port    string
	logger  *log.Logger
	router  *http.ServeMux
	client  *http.Client
	counter int64
	m       appMap // Server should have map, see readme
}

// @title proxxy
// @version 1.0
// @description HTTP server for proxying **HTTP**-requests to 3rd-party services.

// @contact.name doublefint
// @contact.email doublefint@gmail.com

// @host      127.0.0.1:8080
// @BasePath  /

// New application
func New(port string, logger *log.Logger) *App {
	a := App{
		port:   port,
		logger: logger,
		router: http.NewServeMux(),
		client: httpClient(),
		m:      *newAppMap(),
	}
	return &a
}

func (a *App) inc() int64 {
	return atomic.AddInt64(&a.counter, 1)
}

//Addr returns proxxy address
func (a *App) Addr() string {
	return ":" + a.port
}
