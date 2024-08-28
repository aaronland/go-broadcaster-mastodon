package broadcaster

import (
	"context"

	"github.com/aaronland/go-uid"
)

func init() {
	ctx := context.Background()
	RegisterBroadcaster(ctx, "null", NewNullBroadcaster)
}

type NullBroadcaster struct {
	Broadcaster
}

func NewNullBroadcaster(ctx context.Context, uri string) (Broadcaster, error) {

	b := NullBroadcaster{}
	return &b, nil
}

func (b *NullBroadcaster) BroadcastMessage(ctx context.Context, msg *Message) (uid.UID, error) {
	return uid.NewNullUID(ctx)
}
