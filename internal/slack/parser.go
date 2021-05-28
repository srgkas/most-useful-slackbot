package slack

import (
	"fmt"
	"strings"
)

const (
	startTagsDelimiter string = "*Tag(s):*"
	endTagsDelimiter string = "*QA(s):*"
)

type Hotfix struct {
	Project string
	Tag string
}

type HotfixMessage struct {
	// this might be extended with extra fields
	fixes []*Hotfix
}

// GetFixes returns Hotfixes list
func (m *HotfixMessage) GetFixes() []*Hotfix {
	return m.fixes
}

// ParseHotfixMessage parses message from hotfix
func ParseHotfixMessage(message string) (*HotfixMessage, error) {
	fixes, err := extractHotFixes(message)

	if err != nil {
		return nil, err
	}

	if len(fixes) == 0 {
		return nil, fmt.Errorf("no tags found in message")
	}

	return newHotfixMessage(fixes), nil
}

func newHotfixMessage(fixes []*Hotfix) *HotfixMessage {
	return &HotfixMessage{
		fixes: fixes,
	}
}

func newHotFix(project string, tag string) *Hotfix {
	return &Hotfix{
		Project: project,
		Tag: tag,
	}
}

func extractHotFixes(message string) ([]*Hotfix, error) {
	parts := strings.Split(message, "\n")

	var fixes []*Hotfix
	var inTags bool

	for _, part := range parts {
		if part == startTagsDelimiter && !inTags {
			inTags = true
			continue
		}

		if part == endTagsDelimiter {
			break
		}

		if part != endTagsDelimiter && inTags {
			fix, err := tagPartToHotfix(part)

			if err != nil {
				return nil, err
			}

			fixes = append(fixes, fix)
		}
	}

	return fixes, nil
}

func tagPartToHotfix(part string) (*Hotfix, error) {
	link, err := extractLink(part)

	if err != nil {
		return nil, err
	}

	return linkToHotfix(link)
}


func extractLink(text string) (string, error) {
	// parse • <https://link-to-tag-1>
	parts := strings.Split(text, " ")

	if len(parts) != 2 {
		return "", fmt.Errorf("invalid tag list item: %s provided", text)
	}

	cleanLink := strings.ReplaceAll(parts[1], "<", "")
	cleanLink = strings.ReplaceAll(cleanLink, ">", "")

	return cleanLink, nil
}

func linkToHotfix(link string) (*Hotfix, error) {
	// parse https://github.com/srgkas/most-useful-slackbot/releases/tag/v0.0.2
	// to srgkas/most-useful-slackbot and v0.0.2

	domainProjectParts := strings.Split(
		link,
	"https://github.com/",
	)

	if len(domainProjectParts) != 2 {
		return nil, fmt.Errorf("invalid tag link: %s provided", link)
	}

	projectTagParts := strings.Split(
		domainProjectParts[1],
	"/releases/tag/",
	)

	if len(projectTagParts) != 2 {
		return nil, fmt.Errorf("invalid project tag part: %s provided", domainProjectParts[1])
	}

	return newHotFix(projectTagParts[0], projectTagParts[1]), nil
}
