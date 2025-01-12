package handlers

import (
	"context"
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	model_presenter "github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	model_updater "github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/middleware"
)

type router struct {
	presentService presentService
	updaterService updaterService

	multilangManager *lang.Storage
}

type updaterService interface {
	UpdateItemsFromCategory(ctx context.Context, category enums.Category, mode model_updater.CategoryUpdateMode) error
	UpdateTargetItem(ctx context.Context, identifier string, category enums.Category) error

	SubscribeHandler(ctx context.Context, category enums.Category) func(w http.ResponseWriter, r *http.Request)
}

type presentService interface {
	GetItems(ctx context.Context, opt *model_presenter.ApiCardsParams, cat enums.Category) []model_presenter.ItemCartData
	GetUpdates(ctx context.Context, opt *model_presenter.ApiCardsParams, cat enums.Category) []model_presenter.UpdateItem
	GetFilters(ctx context.Context, opt *model_presenter.ApiCardsParams, cat enums.Category) []model_presenter.FilterGroup
}

func New(ctx context.Context, presentService presentService, updaterService updaterService) *router {
	return &router{
		presentService: presentService,
		updaterService: updaterService,

		multilangManager: lang.New(ctx),
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

	HandleFunc("GET /profile",
		r.ProfilePageHandler)
}

func (r *router) InitMuxWithDefaultApi(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) {
	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("PUT /settings", r.SettingsHandler)
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
	HandleFunc("POST /api/updater/{category}", r.RunUpdaterHandler)
	HandleFunc("POST /api/updater/{category}/{identifier}", r.RunItemUpdaterHandler)
	HandleFunc("GET /api/updates/{category}/ws", r.WsUpdaterHandler)
}
