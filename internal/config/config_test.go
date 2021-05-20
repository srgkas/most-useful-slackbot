package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config{}
	config.SetChannels(getChannelsData())
	config.SetDestination(getDestination())
	config.SetToken(getToken())
	config.SetServiceList(getServiceListData())

	assert.Equal(t, config.GetDestination().value, getDestination())
	assert.Equal(t, config.GetChannels().value, getChannelsData())
	assert.Equal(t, config.GetServiceList().value, getServiceListData())
	assert.Equal(t, config.GetToken().value, getToken())
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

func getToken() string {
	return "test_token"
}