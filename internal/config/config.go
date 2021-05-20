package config

type Config struct {
	serviceList ServiceListConf
	destination DestinationConf
	token SlackTokenConf
	channels ChannelsConf
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

type ChannelsConf struct {
	value []string
}

func (c *Config) SetServiceList(values []string) {
	c.serviceList.value = values
}

func (c *Config) SetDestination(value string) {
	c.destination.value = value
}

func (c *Config) SetToken(value string)  {
	c.token.value = value
}

func (c *Config) SetChannels(values []string)  {
	c.channels.value = values
}

func (c *Config) GetServiceList() ServiceListConf {
	return c.serviceList
}

func (c *Config) GetDestination() DestinationConf {
	return c.destination
}

func (c *Config) GetToken() SlackTokenConf {
	return c.token
}

func (c *Config) GetChannels() ChannelsConf {
	return c.channels
}