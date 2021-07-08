package slack

import (
	"fmt"
	"strings"

	slackgo "github.com/slack-go/slack"
)

// Holds user submission handlers

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
}

type SubmissionHandler func(submission *slackgo.InteractionCallback) error

func HandleHotfixSubmission(client *slackgo.Client) SubmissionHandler {
	return func(submission *slackgo.InteractionCallback) error {
		// Interactivity & Shortcuts section to specify url
		// https://api.slack.com/interactivity/handling#setup
		// view_submission
		// See this https://github.com/slack-go/slack/blob/master/examples/modal/modal.go
		// Read data from form and
		// compose message CreateHotfixFormResponseMessage()
		// and post message to thread
		// After submission is received handler should be called in go-routine

		// values are in map by action-id from form
		//fmt.Printf("%+v\n", submission)
		//fmt.Printf("%+v\n", submission.View.State.Values)

		//TODO: At least required validation

		title := submission.View.State.Values["title_block"]["title_input"].Value
		description := submission.View.State.Values["description_block"]["description_input"].Value
		rawServices := submission.View.State.Values["services_block"]["services_input"].Value
		services := strings.Split(rawServices, "\n")
		rawTasks := submission.View.State.Values["tasks_block"]["tasks_input"].Value
		tasks := strings.Split(rawTasks, "\n")
		rawDiffs := submission.View.State.Values["diffs_block"]["diffs_input"].Value
		diffs := strings.Split(rawDiffs, "\n")
		rawTags := submission.View.State.Values["tags_block"]["tags_input"].Value
		tags := strings.Split(rawTags, "\n")
		qas := submission.View.State.Values["qa_block"]["qa_input"].SelectedUsers
		approvers := submission.View.State.Values["approvers_block"]["approvers_input"].SelectedUsers

		// Best interface ever. Why not value?
		channelID := submission.View.State.Values["conv_block"]["conv_input"].SelectedConversation

		hfSubmission := &HotfixSubmission{
			Title:       title,
			Description: description,
			Services:    services,
			Tasks:       tasks,
			Diffs:       diffs,
			Tags:        tags,
			QAs:         qas,
			Approvers:   approvers,
		}

		message := CreateHotfixFormResponseMessage(hfSubmission)

		//TODO: Test with response_url

		//TODO: Why invalid_block soooooqa
		a, b, c, err := client.SendMessage(
			channelID,
			slackgo.MsgOptionBlocks(message.Blocks.BlockSet...),
		)

		if err != nil {
			return err
		}

		fmt.Println(a, b, c)

		return nil
	}
}
