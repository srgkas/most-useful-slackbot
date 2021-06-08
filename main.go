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
		// Interactivity & Shortcuts section to specify url
		// https://api.slack.com/interactivity/handling#setup
		// view_submission
		// See this https://github.com/slack-go/slack/blob/master/examples/modal/modal.go
		var inter slackgo.InteractionCallback
		err := json.Unmarshal([]byte(r.FormValue("payload")), &inter)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// After submission is received handler should be called in go-routine

		// values are in map by action-id from form
		fmt.Printf("%+v\n", inter.View.State.Values)

		w.WriteHeader(200)
		return
	})

	// Handle slash commands
	r.HandleFunc("/commands/hotfix", func (w http.ResponseWriter, r *http.Request) {
		//TODO: Hotfix form handler
		// Create form and publish in hotfix channel
		cmd, err := slackgo.SlashCommandParse(r)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}

		// After slash command is received handler should be called in go-routine

		fmt.Printf("%+v\n", cmd)

		// Naive check or further single mux endpoint
		if cmd.Command != "/hotfix" {
			fmt.Printf("Invalid slash command. Expected /hotfix. Received: %s", cmd.Command)
			w.WriteHeader(400)
			return
		}

		// Move all the bullshittery in separate files for messages
		// This should be composed from json. Probably
		var blocksList []slackgo.Block

		blk := slackgo.NewInputBlock(
			"my-block",
			slackgo.NewTextBlockObject(
				slackgo.PlainTextType,
				"Test title",
				false,
				false,
			),
			slackgo.NewPlainTextInputBlockElement(
				slackgo.NewTextBlockObject(
					slackgo.PlainTextType,
					"Placeholder",
					false,
					false,
				),
				"plain_text_input-action",
			),
		)

		blocksList = append(blocksList, blk)

		modalRequest := slackgo.ModalViewRequest{}
		modalRequest.Type = slackgo.VTModal
		modalRequest.Blocks = slackgo.Blocks{BlockSet: blocksList}
		// Fuck this verbatim shit!!!!
		modalRequest.Title = slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Test title",
			false,
			false,
		)
		modalRequest.Submit = slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Submit",
			false,
			false,
		)
		modalRequest.Close = slackgo.NewTextBlockObject(
			slackgo.PlainTextType,
			"Cancel",
			false,
			false,
		)

		response, _ := slackClient.OpenView(cmd.TriggerID, modalRequest)


		fmt.Printf("%+v\n", response)

		w.WriteHeader(200)
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