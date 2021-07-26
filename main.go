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
var slashCommandHandlers map[string][]slack.SlashCommandHandler
var submissionsHandlers map[string][]slack.SubmissionHandler

var slackClient *slackgo.Client
var githubReleaser gh.Releaser
var cfg *config.Config
var subscriptionsRepo slack.SubscriptionRepo

func main() {
	cfg = config.InitConfig()
	r := mux.NewRouter()

	internal.InitStorage(cfg)
	initSlackClient()
	initGithubReleaser()
	initSubscriptions()
	initHandlers()

	r.HandleFunc("/events/handle", func (w http.ResponseWriter, r *http.Request){
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

	// Handle user form submission
	r.HandleFunc("/interact", func (w http.ResponseWriter, r *http.Request) {
		var inter slackgo.InteractionCallback
		err := json.Unmarshal([]byte(r.FormValue("payload")), &inter)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Printf("%+v\n", inter.View)

		handlerID := inter.View.ExternalID
		handlers, hasHandlers := submissionsHandlers[handlerID]
		if !hasHandlers {
			fmt.Printf("Invalid submission %s received. No handlers found.", handlerID)
			w.WriteHeader(400)
			return
		}

		for _, h := range handlers {
			go func (h slack.SubmissionHandler, submission *slackgo.InteractionCallback) {
				if err := h(submission); err != nil {
					fmt.Printf("Failed to process submission by %v. Error: %v", h, err)
				}
			}(h, &inter)
		}

		w.WriteHeader(200)
		return
	})

	// Handle slash commands
	r.HandleFunc("/commands/hotfix", func (w http.ResponseWriter, r *http.Request) {
		cmd, err := slackgo.SlashCommandParse(r)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}

		handlers, hasHandlers := slashCommandHandlers[cmd.Command]
		if !hasHandlers {
			fmt.Printf("Invalid slash command %s received. No handlers found.", cmd.Command)
			w.WriteHeader(400)
			return
		}

		for _, h := range handlers {
			go func (h slack.SlashCommandHandler, cmd *slackgo.SlashCommand) {
				if err := h(cmd); err != nil {
					fmt.Printf("Failed to process slash command by %v. Error: %v", h, err)
				}
			}(h, &cmd)
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
			internal.Subscribe(subscriptionsRepo),
		},
		"as-deploy-prod": {
			internal.ContainsServiceNameDecorator(
				cfg.GetServiceList().Value,
				internal.Repost(cfg.GetDestinationChannel().Value, slackClient),
			),
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
			//internal.ReleaseTag(githubReleaser, cfg),
			//internal.ParseHotfixMessageExample,
			//internal.Subscribe(subscriptionsRepo),
		},
	}

	slashCommandHandlers = map[string][]slack.SlashCommandHandler{
		slack.HotfixCommand: {
			slack.HotfixCommandHandler(slackClient),
		},
	}

	submissionsHandlers = map[string][]slack.SubmissionHandler{
		slack.HotfixSubmissionModal: {
			slack.HandleHotfixSubmission(slackClient),
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

func initSubscriptions() {
	redis := internal.GetRedisClient()
	subscriptionsRepo = slack.NewSubscriptionRepo(redis)
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