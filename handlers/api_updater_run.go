package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *router) RunUpdaterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ip := r.Header.Get("X-Real-Ip"); ip != "188.68.240.160" {
		http.Error(w, "inacceptable caller", http.StatusNotAcceptable)
		return
	}

	err = s.updaterService.UpdateItemsFromCategory(ctx, cat, model.CategoryUpdateModeWhileNew)
	if _, ok := repository.IsErrorItemNotChanged(err); ok {
		logger.Errorw(ctx, "RunUpdaterHandler attempt to second run", "error", err)
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	logger.Infow(r.Context(), "RunUpdaterHandler completed", "category", cat, "url", r.URL)
}
