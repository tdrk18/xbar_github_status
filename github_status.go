package main

import (
	"time"
)

var (
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

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
	layout := "2006/01/02(Mon) 15:04:05"
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
	case "Minor Service Outage", "Partial System Outage", "Service Under Maintenance":
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
	case "Minor Service Outage", "Partial System Outage", "Service Under Maintenance":
		return "yellow"
	case "Major Service Outage":
		return "red"
	default:
		return "white"
	}
}
