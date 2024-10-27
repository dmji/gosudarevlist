package handlers

import "collector/pkg/recollection/service"

type router struct {
	s service.Service
}

type Pager interface {
}

func New(s service.Service) *router {
	return &router{s: s}
}
