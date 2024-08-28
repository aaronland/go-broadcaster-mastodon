# go-broadcaster-mastodon

Go package implementing the `aaronland/go-broadcaster` interfaces for broadcasting messages to Mastodon.

## Documentation

Documentation is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/broadcast cmd/broadcast/main.go
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
  -verbose
    	Enable verbose (debug) logging.
```

For example:

```
$> ./bin/broadcast \
	-body 'This is a test' \
	-image test.jpg \
	-title 'this is a test' \
	-broadcaster 'mastodon://?credentials={CREDENTIALS}' \
	-verbose

2024/08/27 22:42:49 DEBUG Verbose logging enabled
2024/08/27 22:42:50 DEBUG Upload media for post
2024/08/27 22:42:57 DEBUG Successfully uploaded media id=113038050966671435
2024/08/27 22:42:57 INFO Mastodon post "status ID"=113038051095769392
```

Where `{CREDENTIALS}` is a URL-escaped [sfomuseum/runtimevar](https://github.com/sfomuseum/runtimevar) URI string. For example:

```
constant://?val=oauth2://:{OAUTH2_ACCESSTOKEN}@{MASTODON_HOST}
```

Where `oauth2://:{OAUTH2_ACCESSTOKEN}@{MASTODON_HOST}` is the actual URI string required to create a [go-mastodon-api]() client and the `constant://?val=` is the `runtimevar` part.

This setup is more convoluted than I would like but the "[runtimevar](https://gocloud.dev/howto/runtimevar)" URIs is to allow for tools that post to Mastodon and that required sensitive tokens to be deployed in environments where the management of those tokens can be done in a secure manner. For example, a centralized credentials vault with layered ACLs or even just files on disk with limited permissions.

There is still some work to be done to make it all a bit easier though because URL-escaped strings are always a bit of a nuisance.

## See also

* https://github.com/aaronland/go-broadcaster
* https://github.com/aaronland/go-mastodon-api
* https://github.com/sfomuseum/runtimevar
* https://gocloud.dev/howto/runtimevar