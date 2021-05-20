package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/slack-go/slack"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"io/ioutil"
	"net/http"
)

var slackClient *slack.Client

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
	})

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic("Server can't start")
	}
}

func initSlackClient() {
	conf := config.CreateConfig().GetSlackToken()
	slackClient = slack.New(conf.Value)
}
