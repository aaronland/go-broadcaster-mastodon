package mastodon

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	_ "image"
	"image/jpeg"
	"log/slog"
	"net/url"
	"strconv"
	"time"

	"github.com/aaronland/go-broadcaster"
	"github.com/aaronland/go-mastodon-api/v2/client"
	"github.com/aaronland/go-mastodon-api/v2/response"
	"github.com/aaronland/go-uid"
	"github.com/sfomuseum/runtimevar"
)

func init() {
	ctx := context.Background()
	err := broadcaster.RegisterBroadcaster(ctx, "mastodon", NewMastodonBroadcaster)

	if err != nil {
		panic(err)
	}
}

type MastodonBroadcaster struct {
	broadcaster.Broadcaster
	mastodon_client client.Client
	testing         bool
	dryrun          bool
	quality         int
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

	testing := false
	dryrun := false
	quality := 100

	if q.Has("testing") {

		t, err := strconv.ParseBool(q.Get("testing"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?testing= parameter, %w", err)
		}

		testing = t
	}

	if q.Has("dryrun") {

		d, err := strconv.ParseBool(q.Get("dryrun"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?dryrun= parameter, %w", err)
		}

		dryrun = d
	}

	if q.Has("quality") {

		v, err := strconv.Atoi(q.Get("quality"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?quality= parameter, %w", err)
		}

		quality = v
	}

	br := &MastodonBroadcaster{
		mastodon_client: cl,
		testing:         testing,
		dryrun:          dryrun,
		quality:         quality,
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

			var buf bytes.Buffer
			wr := bufio.NewWriter(&buf)

			// Apparently it's not possible to upload PNG files anymore? That doesn't
			// make any sense but when I try to upload them the Mastodon API returns a
			// 422 Unprocessable Content error

			jpeg_opts := &jpeg.Options{
				Quality: b.quality,
			}

			err := jpeg.Encode(wr, im, jpeg_opts)

			if err != nil {
				return nil, fmt.Errorf("Failed to encode image, %w", err)
			}

			wr.Flush()

			if b.dryrun {
				args.Add("media_ids[]", "dryrun")
			} else {

				br := bytes.NewReader(buf.Bytes())

				slog.Debug("Upload media for post")
				rsp, err := b.mastodon_client.UploadMedia(ctx, br, nil)

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

	slog.Info("Mastodon post successful", "status ID", status_id)
	return uid.NewStringUID(ctx, status_id)
}
