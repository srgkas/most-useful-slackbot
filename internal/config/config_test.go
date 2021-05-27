package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config{}
	config.setChannels(getChannelsData())
	config.setDestinationChannel(getDestination())
	config.setSlackToken(getSlackToken())
	config.setGitToken(getGitToken())
	config.setServiceList(getServiceListData())
	config.setHFApprovalConf(getHFApprovalData())

	assert.Equal(t, config.GetDestinationChannel().Value, getDestination())
	assert.Equal(t, config.GetChannels().Value, getChannelsData())
	assert.Equal(t, config.GetServiceList().Value, getServiceListData())
	assert.Equal(t, config.GetSlackToken().Value, getSlackToken())
	assert.Equal(t, config.GetGitToken().Value, getGitToken())
	assert.Equal(t, config.GetGitToken().Value, getGitToken())
	assert.Equal(t, config.GetHFApprovalConf().Value, getHFApprovalData())
}

func getServiceListData() map[string]ServiceConf {
	return map[string]ServiceConf{
		"first_service": {Github: "first", SearchPhrase: "1st"},
		"second_service": {Github: "second", SearchPhrase: "2nd"},
		"third_service": {Github: "third", SearchPhrase: "3rd"},
	}
}

func getChannelsData() map[string]string {
	return map[string]string{
		"first_channels": "first",
		"second_channels": "second",
		"third_channels": "third",
	}
}

func getDestination() string {
	return "test_destination"
}

func getSlackToken() string {
	return "test_token"
}

func getGitToken() string {
	return "test_git_token"
}

func getHFApprovalData() string {
	return "hf_channel"
}