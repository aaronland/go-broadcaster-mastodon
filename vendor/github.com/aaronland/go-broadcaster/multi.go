package broadcaster

import (
	"context"
	"fmt"
	"github.com/aaronland/go-uid"
	"github.com/hashicorp/go-multierror"
)

type MultiBroadcaster struct {
	Broadcaster
	broadcasters []Broadcaster
	async        bool
}

func NewMultiBroadcasterFromURIs(ctx context.Context, broadcaster_uris ...string) (Broadcaster, error) {

	broadcasters := make([]Broadcaster, len(broadcaster_uris))

	for idx, br_uri := range broadcaster_uris {

		br, err := NewBroadcaster(ctx, br_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to create broadcaster for '%s', %v", br_uri, err)
		}

		broadcasters[idx] = br
	}

	return NewMultiBroadcaster(ctx, broadcasters...)
}

func NewMultiBroadcaster(ctx context.Context, broadcasters ...Broadcaster) (Broadcaster, error) {

	async := true

	b := MultiBroadcaster{
		broadcasters: broadcasters,
		async:        async,
	}

	return &b, nil
}

func (b *MultiBroadcaster) BroadcastMessage(ctx context.Context, msg *Message) (uid.UID, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	th := b.newThrottle()

	done_ch := make(chan bool)
	err_ch := make(chan error)
	id_ch := make(chan uid.UID)

	for _, bc := range b.broadcasters {

		go func(bc Broadcaster, msg *Message) {

			defer func() {
				done_ch <- true
				th <- true
			}()

			<-th

			select {
			case <-ctx.Done():
				return
			default:
				// pass
			}

			id, err := bc.BroadcastMessage(ctx, msg)

			if err != nil {
				err_ch <- fmt.Errorf("[%T] Failed to broadcast message: %s\n", bc, err)
			}

			id_ch <- id

		}(bc, msg)
	}

	remaining := len(b.broadcasters)
	var result error

	ids := make([]uid.UID, 0)

	for remaining > 0 {
		select {
		case <-ctx.Done():
			return uid.NewNullUID(ctx)
		case <-done_ch:
			remaining -= 1
		case err := <-err_ch:
			result = multierror.Append(result, err)
		case id := <-id_ch:
			ids = append(ids, id)
		}
	}

	if result != nil {
		return nil, fmt.Errorf("One or more errors occurred, %w", result)
	}

	return uid.NewMultiUID(ctx, ids...), nil
}

func (b *MultiBroadcaster) newThrottle() chan bool {

	workers := len(b.broadcasters)

	if !b.async {
		workers = 1
	}

	throttle := make(chan bool, workers)

	for i := 0; i < workers; i++ {
		throttle <- true
	}

	return throttle
}
