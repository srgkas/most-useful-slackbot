package slack

import slackgo "github.com/slack-go/slack"

func CreateHotfixFormRequestMessage() slackgo.ModalViewRequest {
	var blocksList []slackgo.Block

	blk := slackgo.NewInputBlock(
		"title_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Test title",
			false,
			false,
		),
		slackgo.NewPlainTextInputBlockElement(
			slackgo.NewTextBlockObject(
				slackgo.PlainTextType,
				"Placeholder",
				false,
				false,
			),
			"title_input",
		),
	)

	converSelect := slackgo.NewOptionsSelectBlockElement(
		slackgo.OptTypeConversations,
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Placeholder",
			false,
			false,
		),
		"conv_input",
	)
	converSelect.DefaultToCurrentConversation = true
	converSelect.ResponseURLEnabled = true

	conver := slackgo.NewInputBlock(
		"conv_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Post to channel",
			false,
			false,
		),
		converSelect,
	)

	blocksList = append(blocksList, blk)
	blocksList = append(blocksList, conver)

	modalRequest := slackgo.ModalViewRequest{}
	modalRequest.CallbackID = "hotfix-modal"

	// Identifies form for submission handler
	modalRequest.ExternalID = "hotfix-modal"
	modalRequest.Type = slackgo.VTModal
	modalRequest.Blocks = slackgo.Blocks{BlockSet: blocksList}
	// Fuck this verbatim shit!!!!
	modalRequest.Title = slackgo.NewTextBlockObject(
		slackgo.PlainTextType,
		"Test title",
		false,
		false,
	)
	modalRequest.Submit = slackgo.NewTextBlockObject(
		slackgo.PlainTextType,
		"Submit",
		false,
		false,
	)
	modalRequest.Close = slackgo.NewTextBlockObject(
		slackgo.PlainTextType,
		"Cancel",
		false,
		false,
	)

	return modalRequest
}

func CreateHotfixFormResponseMessage(s *HotfixSubmission) slackgo.Message {
	var blocks []slackgo.Block

	block := slackgo.NewTextBlockObject(
		slackgo.PlainTextType,
		s.Title,
		false,
		false,
	)
	//TODO: extends blocks with more sections

	blocks = append(blocks, block)

	return slackgo.NewBlockMessage(blocks...)
}