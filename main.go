package main

import (
	"encoding/json"
	"fmt"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"log"
	"os"
)

func CreateConfig() *config.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	services := make(map[string]config.ServiceConf)
	channels := make(map[string]string)
	err = json.Unmarshal([]byte(os.Getenv("SERVICES")), &services)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(os.Getenv("CHANNELS")), &channels)
	if err != nil {
		panic(err)
	}

	c := &config.Config{}
	c.SetServiceList(services)
	c.SetDestinationChannel(os.Getenv("DESTINATION_CHANNEL"))
	c.SetChannels(channels)
	c.SetSlackToken(os.Getenv("SLACK_TOKEN"))
	c.SetGitToken(os.Getenv("GITHUB_TOKEN"))
	return c
}

func main() {
	cfg := CreateConfig()
	fmt.Println(cfg)

	r := mux.NewRouter()

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

	http.ListenAndServe(":8000", r)
}
