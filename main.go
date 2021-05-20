package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	slackgo "github.com/slack-go/slack"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/srgkas/most-useful-slackbot/internal"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/srgkas/most-useful-slackbot/internal/slack"
	"io/ioutil"
	"net/http"
)
var channels map[string]string

var handlersMap = map[string][]internal.Handler {
	"as-hotfixes-approval": {
		internal.Subscribe,
	},
	"as-deploy-prod": {
		internal.Repost,
		internal.ReplyInHotfixThread,
	},
	"as-deploy-prod-au": {
		internal.ReplyInHotfixThread,
	},
	"as-deploy-hf": {
		internal.ReplyInHotfixThread,
	},
}

var slackClient *slackgo.Client

func main() {
	cfg := config.CreateConfig()
	fmt.Println(cfg)

	r := mux.NewRouter()

	initSlackClient()

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

func initSlackClient() {
	conf := config.CreateConfig().GetSlackToken()
	slackClient = slackgo.New(conf.Value)
}

func GetHandlers(e slack.Event) []internal.Handler {
	if ch, ok := channels[e.GetChannel()]; ok {

		if handlers, ok := handlersMap[ch]; ok {
			return handlers
		}
	}

	return []internal.Handler{}
}