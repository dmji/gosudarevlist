package handlers

import (
	"io"
	"net/http"

	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *router) SettingsHandler(w http.ResponseWriter, r *http.Request) {

	bytedata, _ := io.ReadAll(r.Body)
	reqBodyString := string(bytedata)
	logger.Infow(r.Context(), "SettingsHandler body", "url", reqBodyString)

}
