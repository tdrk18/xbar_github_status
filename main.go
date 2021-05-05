package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	ENDPOINT = "https://kctbh9vrtdwd.statuspage.io/api/v2/status.json"
)

var (
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func main() {
	data := getStatus()
	fmt.Println(data.Status.statusEmoji() + " | color=" + data.Status.statusColor())
	fmt.Println("---")
	fmt.Println(data)
	fmt.Println(data.Page.lastUpdatedAt())
}

type GitHubStatus struct {
	Page   Page   `json:"page"`
	Status Status `json:"status"`
}

type Page struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Timezone  string    `json:"time_zone"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (page *Page) lastUpdatedAt() string {
	layout := "2006/01/02 15:04:05"
	return page.UpdatedAt.In(JST).Format(layout)
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

func (status *Status) statusEmoji() string {
	switch status.Description {
	case "All Systems Operational":
		return "✓"
	case "Partial System Outage":
		return "!"
	case "Major Service Outage":
		return "✗"
	default:
		return "?"
	}
}

func (status *Status) statusColor() string {
	switch status.Description {
	case "All Systems Operational":
		return "green"
	case "Partial System Outage":
		return "yellow"
	case "Major Service Outage":
		return "red"
	default:
		return "white"
	}
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
