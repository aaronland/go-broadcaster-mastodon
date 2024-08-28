package main

import (
	"context"
	"log"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/aaronland/go-broadcaster-mastodon"
	"github.com/aaronland/go-broadcaster/app/broadcast"
)

func main() {

	ctx := context.Background()
	err := broadcast.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run broadcast application, %v", err)
	}
}
