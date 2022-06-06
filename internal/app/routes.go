package app

import "context"

func (a *App) routes(ctx context.Context) {
	a.router.HandleFunc("/", a.rootHandler(ctx))
	// TODO: /metrics
	// TODO: /debug
}
