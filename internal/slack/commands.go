package slack

import (
	"fmt"
	slackgo "github.com/slack-go/slack"
)

const HotfixCommand string = "/hotfix"

type SlashCommandHandler func (cmd *slackgo.SlashCommand) error

func HotfixCommandHandler (client *slackgo.Client) SlashCommandHandler {
	return func(cmd *slackgo.SlashCommand) error {
		hotfixModal := CreateHotfixFormRequestMessage()
		response, err := client.OpenView(cmd.TriggerID, hotfixModal)

		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", response)

		return nil
	}
}