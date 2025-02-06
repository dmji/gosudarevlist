package service

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/dmji/gosudarevlist/components/websocket_patches"
	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"
	"github.com/dmji/gosudarevlist/pkg/websocket"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type WsUserData struct{}

type categoryUpdaterData struct {
	ws              *websocket.Manager[WsUserData]
	lastUpdateTimer time.Time
	mx              sync.Mutex
	category        enums.Category
}

func userDataInitializer(ctx context.Context, d *WsUserData) {
}

func (s *categoryUpdaterData) publishUpdate() func(context.Context, io.Writer) error {
	return func(ctx context.Context, w io.Writer) error {
		return websocket_patches.TimerTick([]websocket_patches.Field{
			{
				ClassName: "timer_creted",
				Value:     time_formater.Format(ctx, &s.lastUpdateTimer),
			},
			{
				ClassName: "timer_creted_js",
				Value:     s.lastUpdateTimer.UTC().String(),
			},
			{
				ClassName: "timer_title",
				Value: lang.MustLocalize(ctx,
					&i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "UpdaterPageTimerTitle",
							Other: "{{.Category}} List Updated",
						},
						TemplateData: map[string]string{
							"Category": s.category.Presentation(ctx),
						},
					},
				),
			},
		}).Render(ctx, w)
	}
}
