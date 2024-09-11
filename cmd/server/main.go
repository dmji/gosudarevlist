package main

import (
	"collector/internal/components"
	"collector/internal/handlers"
	"collector/internal/services"
	"collector/pkg/repository"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	repo := repository.New()
	s := services.New(repo)
	r := handlers.New(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HomePageHandler)
	mux.HandleFunc("/anime", r.ShelfPageHandler)
	mux.HandleFunc("/api/cards", r.ApiCards)
	mux.HandleFunc("/parser/animelayer", func(w http.ResponseWriter, r *http.Request) {
		component := components.Page([]components.ScanResult{})
		component.Render(context.Background(), w)
	})
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Println("Server starting on :8080")

	srv := &http.Server{
		Handler: mux,
	}
	conn, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8080))

	log.Fatal(srv.Serve(conn))
}
