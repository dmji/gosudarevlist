package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/components/cards"
	"github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (router *router) LayerItemPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	identifier := r.PathValue("identifier")
	if len(identifier) != 24 {
		logger.Errorw(ctx, "Wrong identifier", "value", identifier)
		http.Error(w, "Wrong identifier: "+identifier, http.StatusInternalServerError)
		return
	}
	cards.CardItemData(&model.ItemCartData{})
}
