package slack

import "encoding/json"

type BlockType struct {
	Type string `json:"type"`
}

type textWithType struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BlockText struct {
	textWithType
	Verbatim bool `json:"verbatim"`
}

type BlockElement struct {
	Type string `json:"type"`
	Elements []textWithType `json:"elements,omitempty"`
}

type BlockSection struct {
	BlockType
	Text BlockText `json:"text,omitempty"`
}

type BlockRichText struct {
	BlockType
	Elements []BlockElement `json:"elements,omitempty"`
}

func (b *BlockSection) GetTexts() []string {
	return []string{b.Text.Text}
}

func (b *BlockRichText) GetTexts() []string {
	var result []string

	for _, e := range b.Elements {
		for _, e := range e.Elements {
			result = append(result, e.Text)
		}
	}

	return result
}

type Auth struct {
	BotID string `json:"bot_id"`
}

type Attachment struct {
	ID     int            `json:"id"`
	Blocks []BlockSection `json:"blocks"`
}

type Message struct {
	Channel         string            `json:"channel"`
	Text            string            `json:"text"`
	Attachments     []Attachment      `json:"attachments,omitempty"`
	Blocks          []json.RawMessage `json:"blocks,omitempty"`
	ThreadTimestamp string            `json:"thread_ts"`
}


type Payload struct {
	Event Event  `json:"event"`
	Type  string `json:"type"`
}

type Event interface {
	GetType() string
	GetText() string
	GetChannel() string
	GetAttachments() []Attachment
	GetBlocks() []json.RawMessage
	GetTimestamp() string
}
