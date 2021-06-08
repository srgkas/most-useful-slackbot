package slack

import slackgo "github.com/slack-go/slack"

const HotfixCommand string = "/hotfix"

func HotfixCommandHandler (client *slackgo.Client, cmd *slackgo.SlashCommand) error {
	// Post form for hotfix
	// CreateHotfixFormRequestMessage
	// Show modal
	return nil
}

// Holds handlers to commands