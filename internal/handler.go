package internal

import (
	"errors"
	slackgo "github.com/slack-go/slack"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/srgkas/most-useful-slackbot/internal/gh"
	"github.com/srgkas/most-useful-slackbot/internal/slack"
	"github.com/srgkas/most-useful-slackbot/internal/teamcity"
)

type Handler func (event slack.Event) error

func Repost(channel string, client *slackgo.Client) Handler {
	return func(event slack.Event) error {
		channelID, err := getChannelId(channel, client)
		if err != nil {
			panic(err)
		}

		messageOptions := slackgo.MsgOptionText(event.Text, false)
		_, _, err = client.PostMessage(channelID, messageOptions)
		return err
	}
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

func getChannelId(channelName string, client *slackgo.Client) (string, error) {
	options := slackgo.GetConversationsParameters{}
	groups, _, err1 := client.GetConversations(&options)
	if err1 != nil {
		panic(err1)
	}

	for _, element := range groups {
		if element.Name == channelName {
			return element.ID, nil
		}
	}

	return "", errors.New("could not find channel with name")
}