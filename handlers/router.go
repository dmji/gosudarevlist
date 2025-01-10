package handlers

import (
	"context"
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/middleware"
)

type router struct {
	presentService   presentService
	multilangManager *lang.Storage
	updaterManagerWs updaterManagerWs
}

type presentService interface {
	GetItems(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.ItemCartData
	GetUpdates(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.UpdateItem
	GetFilters(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.FilterGroup
}

type updaterManagerWs interface {
	SubscribeHandler(w http.ResponseWriter, r *http.Request)
}

func New(ctx context.Context, presentService presentService, updaterManagerWs updaterManagerWs) *router {
	return &router{
		presentService:   presentService,
		multilangManager: lang.New(ctx),
		updaterManagerWs: updaterManagerWs,
	}
}

func (r *router) middlewareHandler(HandleFuncOriginal func(string, func(http.ResponseWriter, *http.Request))) func(string, http.HandlerFunc, ...func(http.HandlerFunc) http.HandlerFunc) {
	return func(pattern string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
		mwstack := []func(http.HandlerFunc) http.HandlerFunc{
			middleware.LangerToContextMiddleware(r.multilangManager),
		}
		n := len(middlewares)
		for i := range n {
			mwstack = append(mwstack, func(h http.HandlerFunc) http.HandlerFunc { return mwstack[i](middlewares[i](h)) })
		}
		HandleFuncOriginal(pattern, mwstack[n](handler))
	}
}

func (r *router) InitMuxWithDefaultPages(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) {
	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("/", r.HomePageHandler)

	HandleFunc("GET /animelayer/{category}",
		r.CollectionListingPageHandler)
	HandleFunc("GET /animelayer/{category}/updates",
		r.CollectionUpdatesPageHandler)

	HandleFunc("GET /animelayer/{category}/updates/ws",
		r.updaterManagerWs.SubscribeHandler)

	HandleFunc("GET /profile",
		r.ProfilePageHandler)
}

func (r *router) InitMuxWithDefaultApi(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) {
	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("GET /api/cards/{category}",
		r.ApiCards,
		middleware.PushQueryFromUrlMiddleware,
	)

	HandleFunc("GET /api/filters/{category}",
		r.ApiFilters,
		middleware.PushQueryFromUrlMiddleware,
		middleware.HxTriggerMiddleware("custom-event-refresh-pages"),
	)

	HandleFunc("GET /api/updates/{category}", r.ApiUpdates)

	HandleFunc("PUT /settings", r.SettingsHandler)
}
