package main

import (
	"fmt"
	"github.com/srgkas/most-useful-slackbot/internal/gh"
)

func main() {
	release := gh.NewRelease("srgkas/most-useful-slackbot", "v0.0.1")
	releaser := gh.CreateReleaser("")

	err := releaser.Release(release)

	if err != nil {
		fmt.Println(err)
	}
}
