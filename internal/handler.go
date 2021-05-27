package internal

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/slack-go/slack"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/srgkas/most-useful-slackbot/internal/gh"
	slackLocal "github.com/srgkas/most-useful-slackbot/internal/slack"
	"github.com/srgkas/most-useful-slackbot/internal/teamcity"
)

var ctx = context.Background()

const DEPLOY_PROD = "as-deploy-prod"
const DEPLOY_PROD_AU  = "as-deploy-prod-au"
const DEPLOY_HF = "as-deploy-hf"

type Handler func (event slack.Event) error

func Repost(event slack.Event) error {
	return nil
}

func Subscribe(event slack.Event) error {
	// subscribe logic
	return nil
}

func ReplyInHotfixThread(
	info teamcity.BuildInfo,
	rdb redis.Client,
	cfg *config.Config,
	sc slack.Client,
	env string) error {

	chVal, err := rdb.Get(ctx, info.Project + "." + info.Tag).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	msText := "Repository: " + info.Project + "\n Tag: " + info.Tag + "\n Env: " + env

	msgOptions := []slack.MsgOption{
		slack.MsgOptionUsername("test"),
		slack.MsgOptionTS(chVal),
		slack.MsgOptionText(msText, false),
	}

	_, _, err = sc.PostMessage(cfg.GetDestinationChannel().Value, msgOptions...)
	if err != nil {
		return err
	}

	return nil
}

func ReleaseTag(releaser gh.Releaser, cfg *config.Config) Handler {
	return func(event slackLocal.Event) error {
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