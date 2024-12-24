package handlers

import (
	"net/http"

	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/middleware"
	"github.com/dmji/gosudarevlist/pkg/recollection/service"
)

type router struct {
	s service.Service
	l *lang.Storage
}

type Pager interface {
}

func New(s service.Service) *router {
	return &router{
		s: s,
		l: lang.New(),
	}
}

func (r *router) middlewareHandler(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) func(pattern string, handler http.HandlerFunc) {
	return func(pattern string, handler http.HandlerFunc) {
		HandleFunOriginal(pattern, middleware.LangerToContextMiddleware(r.l, handler))
	}

}

func (r *router) InitMuxWithDefaultPages(HandleFunOriginal func(string, func(http.ResponseWriter, *http.Request))) {

	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("/", r.HomePageHandler)
	HandleFunc("/animelayer", middleware.HxPushUrlMiddleware(r.CollectionListingPageHandler))
	HandleFunc("/animelayer/updates", middleware.HxPushUrlMiddleware(r.CollectionUpdatesPageHandler))
	HandleFunc("/profile", r.ProfilePageHandler)
}

func (r *router) InitMuxWithDefaultApi(HandleFunOriginal func(pattern string, handler func(http.ResponseWriter, *http.Request))) {

	HandleFunc := r.middlewareHandler(HandleFunOriginal)

	HandleFunc("/api/cards", middleware.HxPushUrlMiddleware(r.ApiCards))
	HandleFunc("/api/filters", middleware.HxPushUrlMiddleware(r.ApiFilters))
	HandleFunc("/api/updates", middleware.HxPushUrlMiddleware(r.ApiUpdates))
}
