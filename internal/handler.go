package internal

import (
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/srgkas/most-useful-slackbot/internal/gh"
	"github.com/srgkas/most-useful-slackbot/internal/slack"
	"github.com/srgkas/most-useful-slackbot/internal/teamcity"
)

type Handler func (event slack.Event) error

func Repost(event slack.Event) error {
	// repost logic
	return nil
}

func Subscribe(event slack.Event) error {
	// subscribe logic
	return nil
}

func ReplyInHotfixThread(event slack.Event) error {
	// reply logic
	return nil
}

func ReleaseTag(releaser gh.Releaser, cfg *config.Config) Handler {
	return func(event slack.Event) error {
		buildInfo, err := teamcity.ParseBuildInfo(event.Text)

		if err != nil {
			return err
		}

		repo, err := buildInfo.GetProjectRepo()

		if err != nil {
			return err
		}

		release := gh.NewRelease(repo, buildInfo.Tag)

		return releaser.Release(release)
	}
}