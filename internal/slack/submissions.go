package slack

import (
	"fmt"
	"strings"

	slackgo "github.com/slack-go/slack"
)

// Holds user submission handlers
type SubmissionHandler func(submission *slackgo.InteractionCallback) error

const HotfixSubmissionModal = "hotfix-modal"

type HotfixSubmission struct {
	Title       string
	Description string
	Services    []string
	Tasks       []string
	Diffs       []string
	Tags        []string
	QAs         []string
	Approvers   []string
	ChannelID   string
}

func HandleHotfixSubmission(client *slackgo.Client) SubmissionHandler {
	return func(submission *slackgo.InteractionCallback) error {
		hfSubmission := createHotFixSubmissionBySubmission(submission)

		message := CreateHotfixFormResponseMessage(hfSubmission)

		a, b, c, err := client.SendMessage(
			hfSubmission.ChannelID,
			slackgo.MsgOptionBlocks(message.Blocks.BlockSet...),
		)

		if err != nil {
			return err
		}

		fmt.Println(a, b, c)

		return nil
	}
}

func createHotFixSubmissionBySubmission(s *slackgo.InteractionCallback) *HotfixSubmission {
	// Interactivity & Shortcuts section to specify url
	// https://api.slack.com/interactivity/handling#setup
	// view_submission
	// See this https://github.com/slack-go/slack/blob/master/examples/modal/modal.go
	// Read data from form and
	// compose message CreateHotfixFormResponseMessage()
	// and post message to thread
	// After submission is received handler should be called in go-routine

	// values are in map by action-id from form
	//fmt.Printf("%+v\n", s)
	//fmt.Printf("%+v\n", s.View.State.Values)

	title := s.View.State.Values["title_block"]["title_input"].Value
	description := s.View.State.Values["description_block"]["description_input"].Value
	rawServices := s.View.State.Values["services_block"]["services_input"].Value
	services := strings.Split(rawServices, "\n")
	rawTasks := s.View.State.Values["tasks_block"]["tasks_input"].Value
	tasks := strings.Split(rawTasks, "\n")
	rawDiffs := s.View.State.Values["diffs_block"]["diffs_input"].Value
	diffs := strings.Split(rawDiffs, "\n")
	rawTags := s.View.State.Values["tags_block"]["tags_input"].Value
	tags := strings.Split(rawTags, "\n")
	qas := s.View.State.Values["qa_block"]["qa_input"].SelectedUsers
	approvers := s.View.State.Values["approvers_block"]["approvers_input"].SelectedUsers
	channelID := s.View.State.Values["conv_block"]["conv_input"].SelectedConversation

	return &HotfixSubmission{
		Title:       title,
		Description: description,
		Services:    services,
		Tasks:       tasks,
		Diffs:       diffs,
		Tags:        tags,
		QAs:         qas,
		Approvers:   approvers,
		ChannelID:   channelID,
	}
}
