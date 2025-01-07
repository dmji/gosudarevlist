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

func (r *router) middlewareHandler(HandleFuncOriginal func(string, func(http.ResponseWriter, *http.Request))) func(string, http.HandlerFunc, ...func(http.HandlerFunc) http.HandlerFunc) {
	return func(pattern string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
		mwstack := []func(http.HandlerFunc) http.HandlerFunc{
			middleware.LangerToContextMiddleware(r.l),
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

	HandleFunc("GET /animelayer/anime",
		r.CollectionListingPageHandler(model.CategoryAnime))
	HandleFunc("GET /animelayer/anime/updates",
		r.CollectionUpdatesPageHandler(model.CategoryAnime))

	HandleFunc("GET /animelayer/anime_hentai",
		r.CollectionListingPageHandler(model.CategoryAnime))
	HandleFunc("GET /animelayer/anime_hentai/updates",
		r.CollectionUpdatesPageHandler(model.CategoryAnime))

	HandleFunc("GET /animelayer/manga",
		r.CollectionListingPageHandler(model.CategoryAnime))
	HandleFunc("GET /animelayer/manga/updates",
		r.CollectionUpdatesPageHandler(model.CategoryAnime))

	HandleFunc("GET /profile",
		r.ProfilePageHandler)
}

func (r *router) InitMuxWithDefaultApi(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) {
	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("GET /api/cards",
		r.ApiCards(model.CategoryAnime),
		middleware.PushQueryFromUrlMiddleware,
	)

	HandleFunc("GET /api/filters",
		r.ApiFilters(model.CategoryAnime),
		middleware.PushQueryFromUrlMiddleware,
		middleware.HxTriggerMiddleware("custom-event-refresh-pages"),
	)

	HandleFunc("GET /api/updates", r.ApiUpdates(model.CategoryAnime))

	HandleFunc("PUT /settings", r.SettingsHandler)
}
