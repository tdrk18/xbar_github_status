package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ENDPOINT = "https://kctbh9vrtdwd.statuspage.io/api/v2/status.json"
)

func getStatus() GitHubStatus {
	req, _ := http.NewRequest("GET", ENDPOINT, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var data GitHubStatus
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
