package internal

import "github.com/srgkas/most-useful-slackbot/internal/slack"

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

func ReleaseTag(event slack.Event) error {
	// String example to parse
	// Succeeded - AirSlate / PROD Env / PROD: Builds / Backend: API / api-addons / Deploy #173 | - v8.11.1 [v8.11.1]>
	return nil
}