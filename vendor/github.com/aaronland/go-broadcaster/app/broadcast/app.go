// Package broadcast provides methods for implementing a command line tool for "broadcasting" messages.
package broadcast

import (
	"context"
	"flag"
	"fmt"
	"image"
	"log/slog"
	"os"

	"github.com/aaronland/go-broadcaster"
	"github.com/sfomuseum/go-flags/flagset"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	br, err := broadcaster.NewMultiBroadcasterFromURIs(ctx, broadcaster_uris...)

	if err != nil {
		return fmt.Errorf("Failed to create broadcaster, %w", err)
	}

	msg := &broadcaster.Message{
		Title: title,
		Body:  body,
	}

	count_images := len(image_paths)

	if count_images > 0 {

		msg.Images = make([]image.Image, count_images)

		for idx, path := range image_paths {

			r, err := os.Open(path)

			if err != nil {
				return fmt.Errorf("Failed to open image %s, %w", path, err)
			}

			defer r.Close()

			im, _, err := image.Decode(r)

			if err != nil {
				return fmt.Errorf("Failed to decode image %s, %w", path, err)
			}

			msg.Images[idx] = im
		}
	}

	id, err := br.BroadcastMessage(ctx, msg)

	if err != nil {
		return fmt.Errorf("Failed to broadcast message, %w", err)
	}

	fmt.Println(id.String())
	return nil
}
