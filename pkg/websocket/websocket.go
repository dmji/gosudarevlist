package websocket

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func NewManager[T comparable](name string, messageBufferSize int, userDataInitializer func(ctx context.Context, d *T)) *Manager[T] {
	return &Manager[T]{
		subscriberMessageBuffer: messageBufferSize,
		subscribers:             make(map[*lang.Loader]map[*subscriber[T]]struct{}),
		name:                    name,
		userDataInitializer:     userDataInitializer,
	}
}

type Manager[T comparable] struct {
	subscriberMessageBuffer int
	subscribersMu           sync.Mutex
	subscribers             map[*lang.Loader]map[*subscriber[T]]struct{}

	name                        string
	fnMsgTextAfterRegisterEvent func(context.Context, *T) []byte
	userDataInitializer         func(ctx context.Context, d *T)
}

func (s *Manager[T]) loggerMsg(body string) string {
	return fmt.Sprintf("WebSocket %s Manager | %s", s.name, body)
}

type subscriber[T comparable] struct {
	userData T
	msgs     chan *message
}

type message struct {
	data        []byte
	messageType websocket.MessageType
}

func (s *Manager[T]) SetAfterRegisterEvent(fn func(context.Context, *T) []byte) {
	s.fnMsgTextAfterRegisterEvent = fn
}

func (s *Manager[T]) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(w, r)
	if err == nil {
		return
	}

	if errors.Is(err, context.Canceled) {
		return
	}

	logger.Errorw(r.Context(), s.loggerMsg("Handler closed with error"), "error", err)
}

func (s *Manager[T]) addSubscriber(ctx context.Context, user *subscriber[T]) {
	langer := lang.FromContext(ctx)
	s.subscribersMu.Lock()
	if _, ok := s.subscribers[langer]; !ok {
		s.subscribers[langer] = make(map[*subscriber[T]]struct{})
	}
	s.subscribers[langer][user] = struct{}{}
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Register subscriber"), "subscriber", user)
}

func (s *Manager[T]) removeSubscriber(ctx context.Context, subscriber *subscriber[T]) {
	langer := lang.FromContext(ctx)
	s.subscribersMu.Lock()
	delete(s.subscribers[langer], subscriber)
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Unregister subscriber"), "subscriber", subscriber)
}

func (s *Manager[T]) subscribe(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var c *websocket.Conn
	subscriber := &subscriber[T]{
		msgs: make(chan *message, s.subscriberMessageBuffer),
	}
	s.userDataInitializer(ctx, &subscriber.userData)
	s.addSubscriber(ctx, subscriber)
	defer s.removeSubscriber(ctx, subscriber)

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
	if s.fnMsgTextAfterRegisterEvent != nil {
		subscriber.msgs <- &message{data: s.fnMsgTextAfterRegisterEvent(ctx, &subscriber.userData), messageType: websocket.MessageText}
	}
	for {
		select {
		case msg := <-subscriber.msgs:
			ctx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()
			err := c.Write(ctx, msg.messageType, msg.data)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (cs *Manager[T]) PublishMsg(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for _, ls := range cs.subscribers {
		for s := range ls {
			s.msgs <- &message{data: msg, messageType: websocket.MessageText}
		}
	}
}

func (cs *Manager[T]) PublishMsgBinary(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for _, ls := range cs.subscribers {
		for s := range ls {
			s.msgs <- &message{data: msg, messageType: websocket.MessageBinary}
		}
	}
}

func (cs *Manager[T]) PublishTempl(Render func(ctx context.Context, w io.Writer) error) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	if len(cs.subscribers) == 0 {
		return
	}

	for l, ls := range cs.subscribers {

		ctx := lang.ToContext(context.Background(), l)
		buf := &bytes.Buffer{}
		Render(ctx, buf)
		b := buf.Bytes()

		for s := range ls {
			s.msgs <- &message{data: b, messageType: websocket.MessageText}
		}
	}
}
