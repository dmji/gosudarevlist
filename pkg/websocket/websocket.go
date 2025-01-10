package websocket

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/dmji/gosudarevlist/lang"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

type subscriberMap map[*subscriber]struct{}

func NewManager(name string, messageBufferSize int) *manager {
	return &manager{
		subscriberMessageBuffer: messageBufferSize,
		subscribers:             make(map[*lang.Loader]map[*subscriber]struct{}),
		name:                    name,
	}
}

type manager struct {
	subscriberMessageBuffer int
	subscribersMu           sync.Mutex
	subscribers             map[*lang.Loader]map[*subscriber]struct{}

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

func (s *manager) addSubscriber(ctx context.Context, user *subscriber) {
	langer := lang.FromContext(ctx)
	s.subscribersMu.Lock()
	if _, ok := s.subscribers[langer]; !ok {
		s.subscribers[langer] = make(map[*subscriber]struct{})
	}
	s.subscribers[langer][user] = struct{}{}
	s.subscribersMu.Unlock()
	logger.Infow(ctx, s.loggerMsg("Register subscriber"), "subscriber", user)
}

func (s *manager) removeSubscriber(ctx context.Context, subscriber *subscriber) {
	langer := lang.FromContext(ctx)
	s.subscribersMu.Lock()
	delete(s.subscribers[langer], subscriber)
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

	for _, ls := range cs.subscribers {
		for s := range ls {
			s.msgs <- &message{data: msg, messageType: websocket.MessageText}
		}
	}
}

func (cs *manager) PublishMsgBinary(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for _, ls := range cs.subscribers {
		for s := range ls {
			s.msgs <- &message{data: msg, messageType: websocket.MessageBinary}
		}
	}
}

func (cs *manager) PublishTempl(Render func(ctx context.Context, w io.Writer) error) {
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
