package slack

import (
	"fmt"
	slackgo "github.com/slack-go/slack"
)

// Holds user submission handlers

const HotfixSubmissionModal = "hotfix-modal"

type HotfixSubmission struct {
	Title string
}

type SubmissionHandler func (submission *slackgo.InteractionCallback) error

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

		title := submission.View.State.Values["title_block"]["title_input"].Value

		// Best interface ever. Why not value?
		channelID := submission.View.State.Values["conv_block"]["conv_input"].SelectedConversation

		//hfSubmission := &HotfixSubmission{
		//	Title: title,
		//}

		//message := CreateHotfixFormResponseMessage(hfSubmission)

		//TODO: Test with response_url


		//TODO: Why invalid_block soooooqa
		a, b, c, err := client.SendMessage(
			channelID,
			//slackgo.MsgOptionBlocks(message.Blocks.BlockSet[0]),
			slackgo.MsgOptionText(title, false),
		)

		if err != nil {
			return err
		}

		fmt.Println(a,b,c)

		return nil
	}
}