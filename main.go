package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageEvent struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	APIAppID string `json:"api_app_id"`
	Event    struct {
		Type        string `json:"type"`
		Channel     string `json:"channel"`
		User        string `json:"user"`
		Text        string `json:"text"`
		Ts          string `json:"ts"`
		EventTs     string `json:"event_ts"`
		ChannelType string `json:"channel_type"`
	} `json:"event"`
	Type        string   `json:"type"`
	AuthedTeams []string `json:"authed_teams"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
}

func main() {
	fmt.Println("vim-go")

	mux := http.NewServeMux()
	mux.HandleFunc("/msg", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request From: ", r.Host)

		var event MessageEvent
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Println(event)

		//json.NewEncoder(w).Encode(&challenge)
	})

	http.ListenAndServe(":8080", mux)
}
