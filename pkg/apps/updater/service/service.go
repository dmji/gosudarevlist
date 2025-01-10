package service

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/dmji/gosudarevlist/components/websocket_patches"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"
	"github.com/dmji/gosudarevlist/pkg/websocket"
)

type ItemProvider interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error)
	GetItemsFromCategoryPages(ctx context.Context, category enums.Category, iPage int) ([]*model.AnimelayerItem, error)
}

type WsUserData struct{}

type service struct {
	repo          repository.AnimeLayerRepositoryDriver
	animelayerApi ItemProvider

	ws              *websocket.Manager[WsUserData]
	lastUpdateTimer *time.Time
}

type Service interface {
	UpdateItemsFromCategory(ctx context.Context, category enums.Category, mode model.CategoryUpdateMode) error
	UpdateTargetItem(ctx context.Context, identifier string, category enums.Category) error
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerApi ItemProvider) *service {
	t := time.Now().Add(-10 * time.Second)
	s := &service{
		repo:            repo,
		animelayerApi:   animelayerApi,
		ws:              websocket.NewManager[WsUserData]("Updater", 10, userDataInitializer),
		lastUpdateTimer: &t,
	}

	s.ws.SetAfterRegisterEvent(
		func(ctx context.Context, _ *WsUserData) []byte {
			buf := &bytes.Buffer{}
			s.publishUpdate()(ctx, buf)
			return buf.Bytes()
		},
	)

	return s
}

func (s *service) WsHandlerProvider() *websocket.Manager[WsUserData] {
	return s.ws
}

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
