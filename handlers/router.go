package handlers

import (
	"context"
	"net/http"

	"github.com/dmji/go-myanimelist/mal/maltype"
	model_presenter "github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/middleware"
)

type router struct {
	presentService presentService
	updaterService updaterService
	mal            malApiService

	multilangManager *lang.Storage
}

type updaterService interface {
	SubscribeHandler(ctx context.Context, category enums.Category) func(w http.ResponseWriter, r *http.Request)
}

type presentService interface {
	GetItems(ctx context.Context, opt *model_presenter.ApiCardsParams, cat enums.Category) []model_presenter.ItemCartData
	GetUpdates(ctx context.Context, opt *model_presenter.ApiCardsParams, cat enums.Category) []model_presenter.UpdateItem
	GetFilters(ctx context.Context, opt *model_presenter.ApiCardsParams, cat enums.Category) []model_presenter.FilterGroup
}
type malApiService interface {
	GetItem(ctx context.Context, id int) (*maltype.Anime, error)
	GetCategory(ctx context.Context) ([]maltype.Anime, error)
}

func New(ctx context.Context, presentService presentService, updaterService updaterService, mal malApiService) *router {
	return &router{
		presentService: presentService,
		updaterService: updaterService,
		mal:            mal,

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

	HandleFunc("GET /animelayer/{category}/{identifier}", r.LayerItemPageHandler)
	HandleFunc("GET /animelayer/{category}", r.CollectionListingPageHandler)
	HandleFunc("GET /animelayer/{category}/updates", r.CollectionUpdatesPageHandler)

	HandleFunc("GET /profile", r.ProfilePageHandler)

	// MAL pages
	HandleFunc("GET /mal/{category}/{id}", r.MalItemPageHandler)
	HandleFunc("GET /mal/{category}", r.MalCategoryPageHandler)

	HandleFunc("/login/{service}", func(w http.ResponseWriter, r *http.Request) {
		s := r.PathValue("service")
		logger.Infow(r.Context(), "Login User Endpoint Reached", "url", r.URL.RawQuery)
		http.Redirect(w, r, "/api/login/"+s, http.StatusAccepted)
	})
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
	// HandleFunc("POST /api/updater/{category}", r.RunUpdaterHandler)
	// HandleFunc("POST /api/updater/{category}/{identifier}", r.RunItemUpdaterHandler)
	HandleFunc("GET /api/updates/{category}/ws", r.WsUpdaterHandler)

	HandleFunc("/api/login/{service}", func(w http.ResponseWriter, r *http.Request) {
		logger.Infow(r.Context(), "Login Api Endpoint Reached", "url", r.URL.RawQuery)
		http.Redirect(w, r, "/profile", http.StatusAccepted)
	})
}
