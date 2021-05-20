package config

type Config struct {
	serviceList        ServiceListConf
	destinationChannel DestinationChannelConf
	slackToken         SlackTokenConf
	gitToken           GitTokenConf
	channels           ChannelsConf
	hfApproval         HFApprovalConf
}

type HFApprovalConf struct {
	value string
}

type ServiceListConf struct {
	value map[string]ServiceConf
}

type ServiceConf struct {
	Github string `json:"github"`
	SearchPhrase string `json:"search-phrase"`
}

type DestinationChannelConf struct {
	value string
}

type SlackTokenConf struct {
	value string
}

type GitTokenConf struct {
	value string
}

type ChannelsConf struct {
	value map[string]string
}

func (c *Config) SetServiceList(values map[string]ServiceConf) {
	c.serviceList.value = values
}

func (c *Config) SetDestinationChannel(value string) {
	c.destinationChannel.value = value
}

func (c *Config) SetSlackToken(value string)  {
	c.slackToken.value = value
}

func (c *Config) SetGitToken(value string)  {
	c.gitToken.value = value
}

func (c *Config) SetChannels(values map[string]string)  {
	c.channels.value = values
}

func (c *Config) SetHFApprovalConf(value string)  {
	c.hfApproval.value = value
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