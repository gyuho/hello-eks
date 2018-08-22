package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "32000"
	}

	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello-eks", handler)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	errc := make(chan error)
	go func() {
		log.Printf("started listener with 'localhost:%s'", port)
		errc <- srv.ListenAndServe()
	}()

	log.Printf("received signal, shutting down server %q", (<-notifier).String())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	srv.Shutdown(ctx)
	cancel()
	log.Printf("shut down server %v", <-errc)

	signal.Stop(notifier)
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Write([]byte(`<b>Hello EKS!</b><br><img src="https://raw.githubusercontent.com/gyuho/hello-eks/master/img/amazon-eks.png" alt="EKS">`))
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
