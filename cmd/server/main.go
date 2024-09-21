package main

import (
	"collector/handlers"
	"collector/internal/services"
	"collector/pkg/middleware"
	repository_inmemory "collector/pkg/repository/inmemory"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	path := ".env"
	for i := range 10 {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}
	panic(".env not found")
}

func main() {

	repo := repository_inmemory.New()
	s := services.New(repo)
	r := handlers.New(s)

	mux := http.NewServeMux()

	// pages
	mux.HandleFunc("/", r.HomePageHandler)
	mux.HandleFunc("/profile", r.ProfilePageHandler)
	mux.HandleFunc("/anime",
		middleware.HxPushUrlMiddleware(r.ShelfPageHandler),
	)

	// parsers
	mux.HandleFunc("/parser/animelayer", r.ScannerPageHandler)

	// api
	mux.HandleFunc("GET /api/push_url",
		middleware.HxPushUrlMiddleware(func(w http.ResponseWriter, r *http.Request) { log.Print("Changed!") }),
	)
	mux.HandleFunc("/api/cards",
		//middleware.PushQueryToRequestMiddleware(
		middleware.HxPushUrlMiddleware(
			r.ApiCards,
		),
		//),
	)

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
