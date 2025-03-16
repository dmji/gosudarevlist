package websocket

import (
	"bytes"
	"context"
	"io"

	"github.com/coder/websocket"
	"github.com/dmji/gosudarevlist/lang"
)

func (cs *manager[T]) PublishMsg(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for _, ls := range cs.subscribers {
		for s := range ls {
			s.msgs <- &message{data: msg, messageType: websocket.MessageText}
		}
	}
}

func (cs *manager[T]) PublishMsgBinary(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for _, ls := range cs.subscribers {
		for s := range ls {
			s.msgs <- &message{data: msg, messageType: websocket.MessageBinary}
		}
	}
}

func (cs *manager[T]) PublishTempl(Render func(ctx context.Context, w io.Writer) error) {
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
