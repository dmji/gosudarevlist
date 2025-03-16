package websocket

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *manager[T]) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(w, r)
	if err == nil {
		return
	}

	if errors.Is(err, context.Canceled) {
		return
	}

	logger.Errorw(r.Context(), s.loggerMsg("Handler closed with error"), "error", err)
}

func (s *manager[T]) addSubscriber(ctx context.Context, user *subscriber[T]) {
	langer := lang.FromContext(ctx)
	s.subscribersMu.Lock()
	if _, ok := s.subscribers[langer]; !ok {
		s.subscribers[langer] = make(map[*subscriber[T]]struct{})
	}
	s.subscribers[langer][user] = struct{}{}
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Register subscriber"), "subscriber", user)
}

func (s *manager[T]) removeSubscriber(ctx context.Context, subscriber *subscriber[T]) {
	langer := lang.FromContext(ctx)
	s.subscribersMu.Lock()
	delete(s.subscribers[langer], subscriber)
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Unregister subscriber"), "subscriber", subscriber)
}

func (s *manager[T]) subscribe(w http.ResponseWriter, r *http.Request) error {
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
