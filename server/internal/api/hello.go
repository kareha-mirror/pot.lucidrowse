package api

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

var messages = []string{
	"おはよう、まだ寝てるの？",
	"こんにちは、ごきげんいかが？",
	"こんばんは、まだ起きてるの？",
}

var messageIndex = 0

func handleHello(w http.ResponseWriter, r *http.Request) {
	res := HelloResponse{
		Message: messages[messageIndex%len(messages)],
	}

	messageIndex++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
