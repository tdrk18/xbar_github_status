package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ENDPOINT = "https://kctbh9vrtdwd.statuspage.io/api/v2/status.json"
)

func main() {
	data := getStatus()
	fmt.Println(data)
}

type GitHubStatus struct {
	Page   Page   `json:"page"`
	Status Status `json:"status"`
}

type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Timezone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

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
