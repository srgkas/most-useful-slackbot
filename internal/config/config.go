package config

type Config struct {
	serviceList ServiceListConf
	destination DestinationConf
	slackToken SlackTokenConf
	gitToken GitTokenConf
	channels ChannelsConf
	hfApproval HFApprovalConf
}

type HFApprovalConf struct {
	value string
}

type ServiceListConf struct {
	value []string
}

type DestinationConf struct {
	value string
}

type SlackTokenConf struct {
	value string
}

type GitTokenConf struct {
	value string
}

type ChannelsConf struct {
	value []string
}

func (c *Config) SetServiceList(values []string) {
	c.serviceList.value = values
}

func (c *Config) SetDestination(value string) {
	c.destination.value = value
}

func (c *Config) SetSlackToken(value string)  {
	c.slackToken.value = value
}

func (c *Config) SetGitToken(value string)  {
	c.gitToken.value = value
}

func (c *Config) SetChannels(values []string)  {
	c.channels.value = values
}

func (c *Config) SetHFApprovalConf(value string)  {
	c.hfApproval.value = value
}

func (c *Config) GetServiceList() ServiceListConf {
	return c.serviceList
}

func (c *Config) GetDestination() DestinationConf {
	return c.destination
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