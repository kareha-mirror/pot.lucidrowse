package api

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Text string `json:"text"`
}

var hellos = []string{
	"おはよう、まだ寝てるの？",
	"こんにちは、ごきげんいかが？",
	"こんばんは、まだ起きてるの？",
}

var helloIndex = 0

func handleHello(w http.ResponseWriter, r *http.Request) {
	res := HelloResponse{
		Text: hellos[helloIndex%len(hellos)],
	}
	helloIndex++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
