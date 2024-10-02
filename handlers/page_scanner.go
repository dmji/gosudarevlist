package handlers

import (
	"collector/components/pages"
	"log"
	"net/http"

	"github.com/dmji/go-animelayer-parser"
)

func (s *router) ScannerPageHandler(w http.ResponseWriter, r *http.Request) {
	//
	log.Print("Handler Scanner | Reached")
	items := make([]animelayer.ItemPartial, 0, 2000)

	log.Printf("Handler Scanner | Items %d", len(items))

	result := make([]pages.ScanResult, 0, len(items))

	result = append(result, pages.ScanResult{
		ID:         1,
		Identifier: "Id",
		Status:     "Ok",
		Title:      "Title",
	})

	log.Printf("Handler Scanner | Items updated %d", len(result))

	err := pages.Scanner(result).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
