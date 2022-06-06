package app

import (
	"context"
	"net/http"
	"time"
)

const (
	srvReadHeaderTimeout = 2
	srvReadTimeout       = 5
	srvWriteTimeout      = 10
	srvIdleTimeout       = 15
	srvShutdownTimeout   = 5
)

//Start starts proxxy
func (a *App) Start(ctx context.Context) error {

	logger := a.logger

	a.routes(ctx)

	srv := &http.Server{
		Addr:              a.Addr(),
		Handler:           a.router,
		ReadHeaderTimeout: srvReadHeaderTimeout * time.Second,
		ReadTimeout:       srvReadTimeout * time.Second,
		WriteTimeout:      srvWriteTimeout * time.Second,
		IdleTimeout:       srvIdleTimeout * time.Second,
		ErrorLog:          logger,
	}

	go func() {

		<-ctx.Done()

		logger.Printf("shutdown...\n")

		downCtx, cancel := context.WithTimeout(context.Background(), srvShutdownTimeout*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(downCtx); err != nil {
			logger.Fatalf("failed to shutdown server: %v\n", err)
		}

	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil

}
