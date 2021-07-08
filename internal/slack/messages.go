package slack

import (
	"strings"

	slackgo "github.com/slack-go/slack"
)

func CreateHotfixFormRequestMessage() slackgo.ModalViewRequest {
	var blocksList []slackgo.Block

	titleBlk := slackgo.NewInputBlock(
		"title_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Title",
			false,
			false,
		),
		slackgo.NewPlainTextInputBlockElement(
			slackgo.NewTextBlockObject(
				slackgo.PlainTextType,
				"Fix title",
				false,
				false,
			),
			"title_input",
		),
	)

	descriptionTextArea := slackgo.NewPlainTextInputBlockElement(
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Fix Description",
			false,
			false,
		),
		"description_input",
	)

	// Multiline indicates textarea
	descriptionTextArea.Multiline = true
	descriptionBlk := slackgo.NewInputBlock(
		"description_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Description",
			false,
			false,
		),
		descriptionTextArea,
	)

	servicesTextArea := slackgo.NewPlainTextInputBlockElement(
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Services (one per line)",
			false,
			false,
		),
		"services_input",
	)

	servicesTextArea.Multiline = true
	servicesBlk := slackgo.NewInputBlock(
		"services_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Service(s)",
			false,
			false,
		),
		servicesTextArea,
	)

	tasksTextArea := slackgo.NewPlainTextInputBlockElement(
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Task links (one per line)",
			false,
			false,
		),
		"tasks_input",
	)

	tasksTextArea.Multiline = true
	tasksBlk := slackgo.NewInputBlock(
		"tasks_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Task(s)",
			false,
			false,
		),
		tasksTextArea,
	)

	diffsTextArea := slackgo.NewPlainTextInputBlockElement(
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Diff links (one per line)",
			false,
			false,
		),
		"diffs_input",
	)

	diffsTextArea.Multiline = true
	diffsBlk := slackgo.NewInputBlock(
		"diffs_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Diffs",
			false,
			false,
		),
		diffsTextArea,
	)

	tagsTextArea := slackgo.NewPlainTextInputBlockElement(
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Tag links (one per line)",
			false,
			false,
		),
		"tags_input",
	)

	tagsTextArea.Multiline = true
	tagsBlk := slackgo.NewInputBlock(
		"tags_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Tag(s)",
			false,
			false,
		),
		tagsTextArea,
	)

	qaSelect := slackgo.NewOptionsMultiSelectBlockElement(
		slackgo.MultiOptTypeUser,
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Select QA",
			false,
			false,
		),
		"qa_input",
	)

	qa := slackgo.NewInputBlock(
		"qa_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"QA(s)",
			false,
			false,
		),
		qaSelect,
	)

	approversSelect := slackgo.NewOptionsMultiSelectBlockElement(
		slackgo.MultiOptTypeUser,
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Select Approvers",
			false,
			false,
		),
		"approvers_input",
	)

	approvers := slackgo.NewInputBlock(
		"approvers_block",
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Approvers",
			false,
			false,
		),
		approversSelect,
	)

	converSelect := slackgo.NewOptionsSelectBlockElement(
		slackgo.OptTypeConversations,
		slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Select conversation",
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

	blocksList = append(blocksList, titleBlk)
	blocksList = append(blocksList, descriptionBlk)
	blocksList = append(blocksList, servicesBlk)
	blocksList = append(blocksList, tasksBlk)
	blocksList = append(blocksList, diffsBlk)
	blocksList = append(blocksList, tagsBlk)
	blocksList = append(blocksList, qa)
	blocksList = append(blocksList, approvers)
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
	var collectedData []string

	title := "*Title:* " + s.Title
	description := "*Description:* " + s.Description
	services := "*Service(s):*\n" + makeListItems(s.Services)
	tasks := "*Task(s):*\n" + makeListItems(s.Tasks)
	diffs := "*Diff(s):*\n" + makeListItems(s.Diffs)
	tags := "*Tag(s):*\n" + makeListItems(s.Tags)
	qa := "*QA(s):*\n" + makeListItems(s.QAs)
	approvers := "*Approvers:*\n" + "* " + strings.Join(s.Approvers, " ")

	collectedData = append(
		collectedData,
		title,
		description,
		services,
		tasks,
		diffs,
		tags,
		qa,
		approvers,
	)

	text := strings.Join(collectedData, "\n")

	block := slackgo.NewSectionBlock(
		slackgo.NewTextBlockObject(
			slackgo.MarkdownType,
			text,
			false,
			false,
		),
		nil,
		nil,
	)

	return slackgo.NewBlockMessage(block)
}

func makeListItems(items []string) string {
	var listItems []string

	for _, item := range items {
		listItems = append(listItems, "* "+item)
	}

	return strings.Join(listItems, "\n")
}
