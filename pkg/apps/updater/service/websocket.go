package service

import (
	"context"
	"io"
	"time"

	"github.com/dmji/gosudarevlist/components/websocket_patches"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"
)

func userDataInitializer(ctx context.Context, d *WsUserData) {
}

func (s *service) runUpdateTicker() {
	for {
		time.Sleep(time.Second)
		s.ws.PublishTempl(s.publishUpdate())
	}
}

func (s *service) publishUpdate() func(context.Context, io.Writer) error {
	return func(ctx context.Context, w io.Writer) error {
		return websocket_patches.TimerTick([]websocket_patches.Field{
			{
				ClassName: "timer_creted",
				Value:     time_formater.Format(ctx, s.lastUpdateTimer),
			},
			{
				ClassName: "timer_creted_js",
				Value:     s.lastUpdateTimer.UTC().String(),
			},
		}).Render(ctx, w)
	}
}
