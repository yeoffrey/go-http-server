package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: Requested /\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Root!\n")
}
func getBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: Requested /blog\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Blog!\n")
}

func serveHttp(ctx context.Context, s *http.Server, ctxCancel context.CancelFunc) {
	err := s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {

		fmt.Printf("server on port %s closed\n", s.Addr)
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
	ctxCancel()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/blog", getBlog)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go serveHttp(ctx, serverOne, cancelCtx)

	<-ctx.Done()
}
