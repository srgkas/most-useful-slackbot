package slack

import "encoding/json"

type Payload struct {
	Token          string
	TeamId         string `json:"team_id"`
	ApiAppId       string `json:"api_app_id"`
	Event          Event  `json:"event"`
	Type           string `json:"type"`
	EventId        string `json:"event_id"`
	EventTime      int    `json:"event_time"`
	Authorizations []struct{
		EnterpriseId string `json:"enterprise_id"`
		TeamId string `json:"team_id"`
		UserId string `json:"user_id"`
		IsBot bool `json:"is_bot"`
		IsEnterpriseInstall bool `json:"is_enterprise_install"`
	}
	IsExtSharedChannel bool `json:"is_ext_shared_channel"`
	EventContext string `json:"event_context"`
}

type Event struct {
	ClientMsgId string `json:"client_msg_id"`
	Type string
	Text string
	User string
	Ts string `json:"ts"`
	Team string
	Blocks []json.RawMessage
	Channel string
	EventTs string
	ChannelType string
}

func (e *Event) GetType() string {
	return e.Type
}

func (e *Event) GetText() string {
	return e.Text
}

func (e *Event) GetChannel() string {
	return e.Channel
}

func (e *Event) GetBlocks() []json.RawMessage {
	return e.Blocks
}

func (e *Event) GetTimestamp() string {
	return e.Ts
}

type EventGetterInterface interface {
	GetType() string
	GetText() string
	GetChannel() string
	GetBlocks() []json.RawMessage
	GetTimestamp() string
}
