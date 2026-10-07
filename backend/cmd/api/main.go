package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/infrastructure/config"
	"github.com/armelo10/vtc_go/backend/internal/interfaces/httpapi"
)

func main() {
	cfg := config.Load()
	server := httpapi.NewServer(cfg)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           server.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("vtc api listening on %s", cfg.HTTPAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	_ = context.Background()
	_ = os.Stdout
}
