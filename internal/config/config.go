package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	serviceList        ServiceListConf
	destinationChannel DestinationChannelConf
	slackToken         SlackTokenConf
	gitToken           GitTokenConf
	channels           ChannelsConf
	hfApproval         HFApprovalConf
	redis              RedisConfig
}

type HFApprovalConf struct {
	Value string
}

type ServiceListConf struct {
	Value map[string]ServiceConf
}

type ServiceConf struct {
	Github       string `json:"github"`
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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DbNumber int
}

func (c *Config) setServiceList(values map[string]ServiceConf) {
	c.serviceList.Value = values
}

func (c *Config) setDestinationChannel(value string) {
	c.destinationChannel.Value = value
}

func (c *Config) setSlackToken(value string) {
	c.slackToken.Value = value
}

func (c *Config) setGitToken(value string) {
	c.gitToken.Value = value
}

func (c *Config) setChannels(values map[string]string) {
	c.channels.Value = values
}

func (c *Config) setHFApprovalConf(value string) {
	c.hfApproval.Value = value
}

func (c *Config) setRedisConf(host, port, password string, db int) {
	c.redis.Host = host
	c.redis.Port = port
}

func (c *Config) GetServiceList() *ServiceListConf {
	return &c.serviceList
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

func (c *Config) GetRedisConf() RedisConfig {
	return c.redis
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		// directly passed env variables will be used instead
		log.Println("Error loading .env file")
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
	c.setServiceList(services)
	c.setDestinationChannel(os.Getenv("DESTINATION_CHANNEL"))
	c.setChannels(channels)
	c.setSlackToken(os.Getenv("SLACK_TOKEN"))
	c.setGitToken(os.Getenv("GITHUB_TOKEN"))
	redisDb, _ := strconv.Atoi((os.Getenv("REDIS_DB")))

	c.setRedisConf(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		redisDb,
	)

	fmt.Println(c)

	return c
}

func GetConfig() *Config {
	return cfg
}