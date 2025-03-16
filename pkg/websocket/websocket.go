package websocket

import (
	"context"
	"fmt"
	"sync"

	"github.com/coder/websocket"
	"github.com/dmji/gosudarevlist/lang"
)

func NewManager[T comparable](name string, messageBufferSize int, userDataInitializer func(ctx context.Context, d *T), fn func(context.Context, *T) []byte) *manager[T] {
	return &manager[T]{
		subscriberMessageBuffer:     messageBufferSize,
		subscribers:                 make(map[*lang.Loader]map[*subscriber[T]]struct{}),
		name:                        name,
		userDataInitializer:         userDataInitializer,
		fnMsgTextAfterRegisterEvent: fn,
	}
}

type subscriber[T comparable] struct {
	userData T
	msgs     chan *message
}

type message struct {
	data        []byte
	messageType websocket.MessageType
}

type manager[T comparable] struct {
	subscriberMessageBuffer int
	subscribersMu           sync.Mutex
	subscribers             map[*lang.Loader]map[*subscriber[T]]struct{}

	name                        string
	fnMsgTextAfterRegisterEvent func(context.Context, *T) []byte
	userDataInitializer         func(ctx context.Context, d *T)
}

func (s *manager[T]) loggerMsg(body string) string {
	return fmt.Sprintf("WebSocket %s Manager | %s", s.name, body)
}
