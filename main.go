package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/events/handle", func (w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8000", r)
}
