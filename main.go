package main

import (
	"fmt"
	"github.com/srgkas/most-useful-slackbot/internal/gh"
)

func main() {
	release := gh.NewRelease("https://github.com/my-repo", "v1.0.0")
	releaser := gh.CreateReleaser("token")

	err := releaser.Release(release)

	if err != nil {
		fmt.Println(err)
	}
}
