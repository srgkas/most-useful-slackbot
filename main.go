package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/srgkas/most-useful-slackbot/internal/slack"
	"io/ioutil"
	"net/http"
)

func main() {
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

	r.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Println("New request from slack message")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var t slack.Payload
		json.Unmarshal(b, &t)

		fmt.Println("Original json:")
		fmt.Println(string(b))

		fmt.Println("Parsed structure:")
		fmt.Printf("%+v\n", t)
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
