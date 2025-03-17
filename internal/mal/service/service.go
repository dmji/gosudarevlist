package service

import (
	"github.com/dmji/go-myanimelist/mal"
)

type service struct {
	client *mal.Site
}

func New(client *mal.Site) *service {
	s := &service{
		client: client,
	}

	return s
}
