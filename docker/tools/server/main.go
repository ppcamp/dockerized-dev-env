package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	_ "embed"
)

//go:embed index.html
var index []byte

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := &http.Server{
		Addr:        ":8080",
		Handler:     routes(),
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()

	slog.Info("serving at http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func routes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", writeHtml(index))
	return r
}

func writeHtml(buffer []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("serving html file")

		if _, err := w.Write(buffer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
}
