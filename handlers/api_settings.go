package handlers

import (
	"io"
	"net/http"

	"github.com/dmji/gosudarevlist/pkg/custom_url"
	"github.com/dmji/gosudarevlist/pkg/expose_header_utils"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (s *router) SettingsHandler(w http.ResponseWriter, r *http.Request) {
	bytedata, _ := io.ReadAll(r.Body)
	reqBodyString := string(bytedata)
	logger.Infow(r.Context(), "SettingsHandler body", "url", reqBodyString)

	settings, err := custom_url.Decode[model.ProfileSettings](reqBodyString)
	if err != nil {
		logger.Errorw(r.Context(), "SettingsHandler | Request decode failed", "string", reqBodyString, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newValue := settings.Language; newValue != nil {
		expose_header_utils.SetCookiePreferedLanguage(w, *newValue)
	}
}
