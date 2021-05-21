package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	slackgo "github.com/slack-go/slack"
	"github.com/srgkas/most-useful-slackbot/internal"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/srgkas/most-useful-slackbot/internal/gh"
	"github.com/srgkas/most-useful-slackbot/internal/slack"
	"io/ioutil"
	"net/http"
)
var channels = make(map[string]string)

var handlersMap map[string][]internal.Handler

var slackClient *slackgo.Client
var githubReleaser gh.Releaser
var cfg *config.Config

func main() {
	cfg = config.CreateConfig()
	fmt.Println(cfg)

	r := mux.NewRouter()

	initSlackClient()
	initGithubReleaser()
	initHandlers()

	r.HandleFunc("/events/handle", func (w http.ResponseWriter, r *http.Request) {
		var err error
		all, _ := ioutil.ReadAll(r.Body)

		// handle challenge request
		var data struct {
			Type string
			Token string
			Challenge string `json:"challenge"`
		}

		fmt.Println(string(all))
		err = json.Unmarshal(all, &data)

		if err == nil && data.Challenge != "" {
			fmt.Println("Challenge accepted")
			w.Write([]byte(data.Challenge))
			return
		}

		// event parsing goes here
		var payload slack.Payload
		json.Unmarshal(all, &payload)
		fmt.Println("Parsed structure:")
		fmt.Printf("%+v\n", payload)

		handlers := GetHandlers(payload.Event)

		for _, h := range handlers {
			go func (h internal.Handler) {
				if err := h(payload.Event); err != nil {
					fmt.Printf("Failed to process event by %v. Error: %v", h, err)
				}
			}(h)
		}
	})

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic("Server can't start")
	}
}

func initHandlers() {
	handlersMap = map[string][]internal.Handler{
		"as-hotfixes-approval": {
			internal.Subscribe,
		},
		"as-deploy-prod": {
			internal.Repost,
			internal.ReplyInHotfixThread,
			internal.ReleaseTag(githubReleaser, cfg),
		},
		"as-deploy-prod-au": {
			internal.ReplyInHotfixThread,
		},
		"as-deploy-hf": {
			internal.ReplyInHotfixThread,
		},
		"silly-willy-test": {
			internal.ReleaseTag(githubReleaser, cfg),
		},
	}
}

func initSlackClient() {
	conf := cfg.GetSlackToken()
	slackClient = slackgo.New(conf.Value)
}

func initGithubReleaser() {
	conf := cfg.GetGitToken()
	githubReleaser = gh.NewReleaser(conf.Value)
}

func GetHandlers(e slack.Event) []internal.Handler {
	// We need default channels lookup.
	// For now we get channels one by one and store in run-time cache.
	channelID := e.GetChannel()

	channelName, ok := channels[channelID]

	if !ok {
		channel, err := slackClient.GetConversationInfo(channelID, false)
		if err != nil {
			fmt.Println(err)
			return []internal.Handler{}
		}

		// Naive cache
		channelName = channel.Name
		channels[channelID] = channelName
	}

	if handlers, ok := handlersMap[channelName]; ok {
		fmt.Println("Got handlers")
		return handlers
	}

	fmt.Println("No handlers")

	return []internal.Handler{}
}