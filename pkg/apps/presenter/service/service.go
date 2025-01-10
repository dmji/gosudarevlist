package service

import (
	"context"
	"io"
	"time"

	"github.com/dmji/gosudarevlist/components/websocket_patches"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"
)

type wsPublishet interface {
	PublishTempl(Render func(ctx context.Context, w io.Writer) error)
}

type service struct {
	repository.AnimeLayerRepositoryDriver
	ws wsPublishet
}

func New(repo repository.AnimeLayerRepositoryDriver, ws wsPublishet) *service {
	s := &service{
		AnimeLayerRepositoryDriver: repo,
		ws:                         ws,
	}
	go s.runUpdateTicker()
	return s
}

func (s *service) runUpdateTicker() {
	timer := time.Now()
	for {
		time.Sleep(time.Second)
		s.ws.PublishTempl(func(ctx context.Context, w io.Writer) error {
			return websocket_patches.TimerTick([]websocket_patches.Field{
				{"timer_creted", time_formater.Format(ctx, &timer)},
				{"timer_creted_js", timer.Format("2006-01-02:15:04:05")},
			}).Render(ctx, w)
		})
	}
}
