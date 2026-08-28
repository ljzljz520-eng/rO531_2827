package main

import (
	portal "hospitalportal"
	"hospitalportal/internal/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	dbPath := os.Getenv("PORTAL_DB")
	if dbPath == "" {
		dbPath = "data/portal.db"
	}
	p, err := portal.Open(dbPath, time.Now)
	if err != nil {
		slog.Error("open portal", "error", err)
		os.Exit(1)
	}
	defer p.Close()
	server := api.New(p, slog.Default())
	address := os.Getenv("PORTAL_ADDR")
	if address == "" {
		address = ":8080"
	}
	slog.Info("portal listening", "address", address)
	if err = http.ListenAndServe(address, server.Handler()); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
