# go-broadcaster

Minimalist and opinionated package providing interfaces for "broadcasting" messages with zero or more images.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-broadcaster.svg)](https://pkg.go.dev/github.com/aaronland/go-broadcaster)

## Motivation

This package provides minimalist and opinionated interfaces for "broadcasting" simple messages with zero or more images.

A message consists of an optional title and body as well as zero or more images. How those elements are processed is left to service or target -specific implementations of the `Broadcaster` interfaces.

That's all it does and doesn't try to account for any other, or more complicated, publishing scenarios. There are other tools and packages for that.

## This should still be considered "in flux"

Although the "skeleton" of this package and its interfaces is complete some details may still change. For example, it is expected that the `Broadcaster` interface will be updated to include a `Close` method shortly.

## Example

### Broadcasting a message

```
package broadcast

import (
	"context"
	"fmt"
	
	"github.com/aaronland/go-broadcaster"
)

func main() {

	ctx := context.Background()     

	br, _ := broadcaster.NewBroadcaster(ctx, "log://")

	msg := &broadcaster.Message{
		Title: "This is the title",
		Body:  "This is a message",
	}

	id, _ := br.BroadcastMessage(ctx, msg)

	fmt.Println(id.String())
	return nil
}
```

_Error handling omitted for the sake of brevity._

### Implementing the `Broadcaster` interface.

```
package broadcaster

import (
	"context"
	"github.com/aaronland/go-uid"
	"log"
	"time"
)

func init() {
	ctx := context.Background()
	RegisterBroadcaster(ctx, "log", NewLogBroadcaster)
}

// LogBroadcaster implements the `Broadcaster` interface to broadcast messages
type LogBroadcaster struct {
	Broadcaster
}

// NewLogBroadcaster returns a new `LogBroadcaster` configured by 'uri' which is expected to
// take the form of:
//
//	log://
//
// By default `LogBroadcaster` instances are configured to broadcast messages to a `log.Default`
// instance. If you want to change that call the `SetLogger` method.
func NewLogBroadcaster(ctx context.Context, uri string) (Broadcaster, error) {
	
	logger := log.Default()
	
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
```

## UIDs

This package uses the [aaronland/go-uid](https://github.com/aaronland/go-uid) package to encapsulate unique identifiers returns by broadcasting targets.

## Tools

```
$> make cli
go build -mod vendor -o bin/broadcast cmd/broadcast/main.go
```

### broadcast

```
$> ./bin/broadcast -h
  -body string
    	The body of the message to broadcast.
  -broadcaster value
    	One or more aaronland/go-broadcast URIs.
  -image value
    	Zero or more paths to images to include with the message to broadcast.
  -title string
    	The title of the message to broadcast.
```

For example:

```
$> ./bin/broadcast -broadcaster log:// -broadcaster null:// -body "hello world"
2022/11/08 08:44:47  hello world
NullUID# Int64UID#1667925887
```

## Other implementations

### flickr://

* https://github.com/aaronland/go-broadcaster-flickr

### mastodon://

* https://github.com/aaronland/go-broadcaster-mastodon

### slack://

* https://github.com/aaronland/go-broadcaster-slack

### twitter://

* https://github.com/aaronland/go-broadcaster-twitter

## See also

* https://github.com/aaronland/go-uid
