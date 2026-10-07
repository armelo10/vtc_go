package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/infrastructure/config"
	"github.com/armelo10/vtc_go/backend/internal/infrastructure/postgres"
	"github.com/armelo10/vtc_go/backend/internal/interfaces/httpapi"
)

func main() {
	cfg := config.Load()
	pool, err := postgres.Open(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewServer(cfg, pool).Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("vtc api listening on %s", cfg.HTTPAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
