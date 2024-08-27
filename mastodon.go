package mastodon

import (
	"bytes"
	"context"
	"fmt"
	_ "image"
	"log/slog"
	"net/url"
	"strconv"
	"time"

	"github.com/aaronland/go-broadcaster"
	"github.com/aaronland/go-image-encode"
	"github.com/aaronland/go-mastodon-api/v2/client"
	"github.com/aaronland/go-mastodon-api/v2/response"
	"github.com/aaronland/go-uid"
	"github.com/sfomuseum/runtimevar"
)

func init() {
	ctx := context.Background()
	broadcaster.RegisterBroadcaster(ctx, "mastodon", NewMastodonBroadcaster)
}

type MastodonBroadcaster struct {
	broadcaster.Broadcaster
	mastodon_client client.Client
	testing         bool
	dryrun          bool
	encoder         encode.Encoder
}

func NewMastodonBroadcaster(ctx context.Context, uri string) (broadcaster.Broadcaster, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	creds_uri := q.Get("credentials")

	if creds_uri == "" {
		return nil, fmt.Errorf("Missing ?credentials= parameter")
	}

	rt_ctx, rt_cancel := context.WithTimeout(ctx, 5*time.Second)
	defer rt_cancel()

	client_uri, err := runtimevar.StringVar(rt_ctx, creds_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive URI from credentials, %w", err)
	}

	cl, err := client.NewClient(ctx, client_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new Mastodon client, %w", err)
	}

	enc, err := encode.NewEncoder(ctx, "png://")

	if err != nil {
		return nil, fmt.Errorf("Failed to create image encoder, %w", err)
	}

	testing := false
	dryrun := false

	str_testing := q.Get("testing")

	if str_testing != "" {

		t, err := strconv.ParseBool(str_testing)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?testing= parameter, %w", err)
		}

		testing = t
	}

	str_dryrun := q.Get("dryrun")

	if str_dryrun != "" {

		d, err := strconv.ParseBool(str_dryrun)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?dryrun= parameter, %w", err)
		}

		dryrun = d
	}

	br := &MastodonBroadcaster{
		mastodon_client: cl,
		testing:         testing,
		dryrun:          dryrun,
		encoder:         enc,
	}

	return br, nil
}

func (b *MastodonBroadcaster) BroadcastMessage(ctx context.Context, msg *broadcaster.Message) (uid.UID, error) {

	status := msg.Body

	if b.testing {
		status = fmt.Sprintf("this is a test and there may be more / please disregard and apologies for the distraction / meanwhile: %s", status)
	}

	args := &url.Values{}

	args.Set("status", status)
	args.Set("visibility", "public")

	if len(msg.Images) > 0 {

		for _, im := range msg.Images {

			// but what if GIF...

			r := new(bytes.Buffer)

			err := b.encoder.Encode(ctx, im, r)

			if err != nil {
				return nil, fmt.Errorf("Failed to encode image, %w", err)
			}

			if b.dryrun {
				args.Add("media_ids[]", "dryrun")
			} else {

				slog.Debug("Upload media for post")
				rsp, err := b.mastodon_client.UploadMedia(ctx, r, nil)

				if err != nil {
					return nil, fmt.Errorf("Failed to upload image, %w", err)
				}

				media_id, err := response.Id(ctx, rsp)

				if err != nil {
					return nil, fmt.Errorf("Failed to derive media ID from response, %w", err)
				}

				slog.Debug("Successfully uploaded media", "id", media_id)
				args.Add("media_ids[]", media_id)
			}
		}
	}

	var status_id string

	if b.dryrun {
		slog.Info("Dryrun", "args", args)
		status_id = "1"
	} else {

		rsp, err := b.mastodon_client.ExecuteMethod(ctx, "POST", "/api/v1/statuses", args)

		if err != nil {
			return nil, fmt.Errorf("Failed to post message, %w", err)
		}

		id, err := response.Id(ctx, rsp)

		if err != nil {
			return nil, fmt.Errorf("Failed to derive status ID from response, %w", err)
		}

		status_id = id
	}

	slog.Info("Mastodon post", "status ID", status_id)

	return uid.NewStringUID(ctx, status_id)
}
