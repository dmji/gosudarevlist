package websocket

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func NewManager(name string, messageBufferSize int) *manager {
	return &manager{
		subscriberMessageBuffer: messageBufferSize,
		subscribers:             make(map[*subscriber]struct{}),
		name:                    name,
	}
}

type manager struct {
	subscriberMessageBuffer int
	subscribersMu           sync.Mutex
	subscribers             map[*subscriber]struct{}

	name string
}

func (s *manager) loggerMsg(body string) string {
	return fmt.Sprintf("WebSocket %s Manager | %s", s.name, body)
}

type subscriber struct {
	msgs chan *message
}

type message struct {
	data        []byte
	messageType websocket.MessageType
}

func (s *manager) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(w, r)
	if err != nil {
		logger.Errorw(r.Context(), s.loggerMsg("Handler closed with error"), "error", err)
		return
	}
}

func (s *manager) addSubscriber(ctx context.Context, subscriber *subscriber) {
	s.subscribersMu.Lock()
	s.subscribers[subscriber] = struct{}{}
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Register subscriber"), "subscriber", subscriber)
}

func (s *manager) removeSubscriber(ctx context.Context, subscriber *subscriber) {
	s.subscribersMu.Lock()
	delete(s.subscribers, subscriber)
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Unregister subscriber"), "subscriber", subscriber)
}

func (s *manager) subscribe(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var c *websocket.Conn
	subscriber := &subscriber{
		msgs: make(chan *message, s.subscriberMessageBuffer),
	}
	s.addSubscriber(ctx, subscriber)
	defer s.removeSubscriber(ctx, subscriber)

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
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

func (cs *manager) PublishMsg(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for s := range cs.subscribers {
		s.msgs <- &message{data: msg, messageType: websocket.MessageText}
	}
}

func (cs *manager) PublishMsgBinary(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for s := range cs.subscribers {
		s.msgs <- &message{data: msg, messageType: websocket.MessageBinary}
	}
}
