package handlers

/*
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

func (s *router) RunItemUpdaterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cat, err := enums.CategoryFromString(r.PathValue("category"))
	if err != nil {
		logger.Errorw(ctx, "PathValue parsing failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	identifier := r.PathValue("identifier")
	if len(identifier) == 0 {
		e := "identifier should be not empty"
		logger.Errorw(ctx, e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	if ip := r.Header.Get("X-Real-Ip"); ip != "188.68.240.160" {
		http.Error(w, "inacceptable caller", http.StatusNotAcceptable)
		return
	}

	err = s.updaterService.UpdateTargetItem(ctx, identifier, cat)
	if _, ok := repository.IsErrorItemNotChanged(err); ok {
		logger.Errorw(ctx, "RunItemUpdaterHandler attempt to second run", "error", err)
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	logger.Infow(r.Context(), "RunItemUpdaterHandler completed", "category", cat, "identifier", identifier, "url", r.URL)
}
*/
