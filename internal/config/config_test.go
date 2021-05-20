package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config{}
	config.SetChannels(getChannelsData())
	config.SetDestination(getDestination())
	config.SetSlackToken(getSlackToken())
	config.SetGitToken(getGitToken())
	config.SetServiceList(getServiceListData())
	config.SetHFApprovalConf(getHFApprovalData())

	assert.Equal(t, config.GetDestination().value, getDestination())
	assert.Equal(t, config.GetChannels().value, getChannelsData())
	assert.Equal(t, config.GetServiceList().value, getServiceListData())
	assert.Equal(t, config.GetSlackToken().value, getSlackToken())
	assert.Equal(t, config.GetGitToken().value, getGitToken())
	assert.Equal(t, config.GetGitToken().value, getGitToken())
	assert.Equal(t, config.GetHFApprovalConf().value, getHFApprovalData())
}

func getServiceListData() []string {
	return []string{"first_service", "second_service", "third_service"}
}

func getChannelsData() []string {
	return []string{"first_channels", "second_channels", "third_channels"}
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