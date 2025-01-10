package service

import (
	"bytes"
	"context"
	"time"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
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
