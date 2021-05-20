package main

import (
	"encoding/json"
	"fmt"
	"github.com/srgkas/most-useful-slackbot/internal/config"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func main() {
	cfg := config.CreateConfig()
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
