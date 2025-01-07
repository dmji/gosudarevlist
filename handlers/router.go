package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/middleware"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	"github.com/dmji/gosudarevlist/pkg/recollection/service"
)

type router struct {
	s service.Service
	l *lang.Storage
}

type Pager interface{}

func New(s service.Service) *router {
	return &router{
		s: s,
		l: lang.New(),
	}
}

func (r *router) middlewareHandler(HandleFuncOriginal func(string, func(http.ResponseWriter, *http.Request))) func(pattern string, handler http.HandlerFunc) {
	return func(pattern string, handler http.HandlerFunc) {
		HandleFuncOriginal(pattern, middleware.LangerToContextMiddleware(r.l, handler))
	}
}

func (r *router) InitMuxWithDefaultPages(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) {
	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("/", r.HomePageHandler)

	HandleFunc("GET /animelayer/anime", middleware.HxPushUrlMiddleware(r.CollectionListingPageHandler(model.CategoryAnime)))
	HandleFunc("GET /animelayer/anime/updates", middleware.HxPushUrlMiddleware(r.CollectionUpdatesPageHandler(model.CategoryAnime)))

	HandleFunc("GET /animelayer/anime_hentai", middleware.HxPushUrlMiddleware(r.CollectionListingPageHandler(model.CategoryAnime)))
	HandleFunc("GET /animelayer/anime_hentai/updates", middleware.HxPushUrlMiddleware(r.CollectionUpdatesPageHandler(model.CategoryAnime)))

	HandleFunc("GET /animelayer/manga", middleware.HxPushUrlMiddleware(r.CollectionListingPageHandler(model.CategoryAnime)))
	HandleFunc("GET /animelayer/manga/updates", middleware.HxPushUrlMiddleware(r.CollectionUpdatesPageHandler(model.CategoryAnime)))

	HandleFunc("GET /profile", r.ProfilePageHandler)
}

func (r *router) InitMuxWithDefaultApi(HandleFunOriginal func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("GET /api/cards", r.ApiCards(model.CategoryAnime))
	HandleFunc("GET /api/filters", middleware.HxPushUrlMiddleware(r.ApiFilters(model.CategoryAnime)))
	HandleFunc("GET /api/updates", r.ApiUpdates(model.CategoryAnime))

	HandleFunc("PUT /settings", r.SettingsHandler)
}
