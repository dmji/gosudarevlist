package main

import (
	"collector/components/pages"
	"collector/internal/handlers"
	"collector/internal/services"
	repository_inmemory "collector/pkg/repository/inmemory"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {

	repo := repository_inmemory.New()
	s := services.New(repo)
	r := handlers.New(s)

	mux := http.NewServeMux()

	// pages
	mux.HandleFunc("/", r.HomePageHandler)
	mux.HandleFunc("/profile", r.ProfilePageHandler)
	mux.HandleFunc("/anime", r.ShelfPageHandler)

	// parsers
	mux.HandleFunc("/parser/animelayer", func(w http.ResponseWriter, r *http.Request) {
		component := pages.Scanner([]pages.ScanResult{})
		component.Render(context.Background(), w)
	})

	// api
	mux.HandleFunc("/api/cards", r.ApiCards)
	mux.HandleFunc("/api/parser/animelayer/category", r.ApiMyAnimeListParseCategory)
	mux.HandleFunc("/api/parser/animelayer/page", r.ApiMyAnimeListParsePage)

	// static assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// starting
	log.Println("Server starting on :8080")

	srv := &http.Server{
		Handler: mux,
	}
	conn, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8080))

	log.Fatal(srv.Serve(conn))
}
