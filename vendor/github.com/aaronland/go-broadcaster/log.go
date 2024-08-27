package broadcaster

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/aaronland/go-uid"
)

func init() {
	ctx := context.Background()
	RegisterBroadcaster(ctx, "log", NewLogBroadcaster)
}

// LogBroadcaster implements the `Broadcaster` interface to broadcast messages
// to a `log.Logger` instance.
type LogBroadcaster struct {
	Broadcaster
}

// NewLogBroadcaster returns a new `LogBroadcaster` configured by 'uri' which is expected to
// take the form of:
//
//	log://
//
// By default `LogBroadcaster` instances are configured to broadcast messages to a `log/slog.Default`
// instance with an `INFO` level.
func NewLogBroadcaster(ctx context.Context, uri string) (Broadcaster, error) {

	b := LogBroadcaster{}
	return &b, nil
}

// BroadcastMessage broadcast the title and body properties of 'msg' to the `log.Logger` instance
// associated with 'b'. It does not publish images yet. Maybe someday it will try to convert images
// to their ascii interpretations but today it does not. It returns the value of the Unix timestamp
// that the log message was broadcast.
func (b *LogBroadcaster) BroadcastMessage(ctx context.Context, msg *Message) (uid.UID, error) {

	log_msg := fmt.Sprintf("%s %s", msg.Title, msg.Body)

	logger := slog.Default()
	logger.Info(log_msg)

	now := time.Now()
	ts := now.Unix()

	return uid.NewInt64UID(ctx, ts)
}
