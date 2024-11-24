package handlers

import (
	"collector/pkg/middleware"
	"collector/pkg/recollection/service"
	"net/http"
)

type router struct {
	s service.Service
}

type Pager interface {
}

func New(s service.Service) *router {
	return &router{s: s}
}

func (r *router) InitMuxWithDefaultPages(HandleFunc func(pattern string, handler func(http.ResponseWriter, *http.Request))) {

	HandleFunc("/", r.HomePageHandler)

	HandleFunc("/animelayer", middleware.HxPushUrlMiddleware(r.ShelfPageHandler))
	HandleFunc("/animelayer/updates", r.UpdatesListHandler)

	HandleFunc("/profile", r.ProfilePageHandler)
}

func (r *router) InitMuxWithDefaultApi(HandleFunc func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	HandleFunc("/api/cards", middleware.HxPushUrlMiddleware(r.ApiCards))
}
