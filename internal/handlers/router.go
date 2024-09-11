package handlers

import "collector/internal/services"

type router struct {
	s services.Service
}

type Pager interface {
}

func New(s services.Service) *router {
	return &router{s: s}
}
