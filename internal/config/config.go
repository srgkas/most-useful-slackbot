package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	serviceList        ServiceListConf
	destinationChannel DestinationChannelConf
	slackToken         SlackTokenConf
	gitToken           GitTokenConf
	channels           ChannelsConf
	hfApproval         HFApprovalConf
}

type HFApprovalConf struct {
	Value string
}

type ServiceListConf struct {
	Value map[string]ServiceConf
}

type ServiceConf struct {
	Github string `json:"github"`
	SearchPhrase string `json:"search-phrase"`
}

type DestinationChannelConf struct {
	Value string
}

type SlackTokenConf struct {
	Value string
}

type GitTokenConf struct {
	Value string
}

type ChannelsConf struct {
	Value map[string]string
}

func (c *Config) SetServiceList(values map[string]ServiceConf) {
	c.serviceList.Value = values
}

func (c *Config) SetDestinationChannel(value string) {
	c.destinationChannel.Value = value
}

func (c *Config) SetSlackToken(value string)  {
	c.slackToken.Value = value
}

func (c *Config) SetGitToken(value string)  {
	c.gitToken.Value = value
}

func (c *Config) SetChannels(values map[string]string)  {
	c.channels.Value = values
}

func (c *Config) SetHFApprovalConf(value string)  {
	c.hfApproval.Value = value
}

func (c *Config) GetServiceList() ServiceListConf {
	return c.serviceList
}

func (c *Config) GetDestinationChannel() DestinationChannelConf {
	return c.destinationChannel
}

func (c *Config) GetSlackToken() SlackTokenConf {
	return c.slackToken
}

func (c *Config) GetGitToken() GitTokenConf {
	return c.gitToken
}

func (c *Config) GetChannels() ChannelsConf {
	return c.channels
}

func (c *Config) GetHFApprovalConf() HFApprovalConf {
	return c.hfApproval
}

func CreateConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	services := make(map[string]ServiceConf)
	channels := make(map[string]string)
	err = json.Unmarshal([]byte(os.Getenv("SERVICES")), &services)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(os.Getenv("CHANNELS")), &channels)
	if err != nil {
		panic(err)
	}

	c := &Config{}
	c.SetServiceList(services)
	c.SetDestinationChannel(os.Getenv("DESTINATION_CHANNEL"))
	c.SetChannels(channels)
	c.SetSlackToken(os.Getenv("SLACK_TOKEN"))
	c.SetGitToken(os.Getenv("GITHUB_TOKEN"))
	return c
}