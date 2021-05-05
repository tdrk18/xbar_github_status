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
	ICON     = "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAKRlWElmTU0AKgAAAAgABQESAAMAAAABAAEAAAEaAAUAAAABAAAASgEbAAUAAAABAAAAUgExAAIAAAAgAAAAWodpAAQAAAABAAAAegAAAAAAAABIAAAAAQAAAEgAAAABQWRvYmUgUGhvdG9zaG9wIENTNiAoTWFjaW50b3NoKQAAA6ABAAMAAAABAAEAAKACAAQAAAABAAAAEKADAAQAAAABAAAAEAAAAAAXm6VPAAAACXBIWXMAAAsTAAALEwEAmpwYAAAEE2lUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNi4wLjAiPgogICA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPgogICAgICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIgogICAgICAgICAgICB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIKICAgICAgICAgICAgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiCiAgICAgICAgICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iPgogICAgICAgICA8eG1wTU06RG9jdW1lbnRJRD54bXAuZGlkOkREQjFCMDlGODZDRTExRTNBQTUyRUUzMzUyRDFCQzQ2PC94bXBNTTpEb2N1bWVudElEPgogICAgICAgICA8eG1wTU06SW5zdGFuY2VJRD54bXAuaWlkOkREQjFCMDlFODZDRTExRTNBQTUyRUUzMzUyRDFCQzQ2PC94bXBNTTpJbnN0YW5jZUlEPgogICAgICAgICA8eG1wTU06RGVyaXZlZEZyb20gcmRmOnBhcnNlVHlwZT0iUmVzb3VyY2UiPgogICAgICAgICAgICA8c3RSZWY6aW5zdGFuY2VJRD54bXAuaWlkOkU1MTc4QTJBOTlBMDExRTI5QTE1QkMxMDQ2QTg5MDREPC9zdFJlZjppbnN0YW5jZUlEPgogICAgICAgICAgICA8c3RSZWY6ZG9jdW1lbnRJRD54bXAuZGlkOkU1MTc4QTJCOTlBMDExRTI5QTE1QkMxMDQ2QTg5MDREPC9zdFJlZjpkb2N1bWVudElEPgogICAgICAgICA8L3htcE1NOkRlcml2ZWRGcm9tPgogICAgICAgICA8eG1wOkNyZWF0b3JUb29sPkFkb2JlIFBob3Rvc2hvcCBDUzYgKE1hY2ludG9zaCk8L3htcDpDcmVhdG9yVG9vbD4KICAgICAgICAgPHRpZmY6T3JpZW50YXRpb24+MTwvdGlmZjpPcmllbnRhdGlvbj4KICAgICAgPC9yZGY6RGVzY3JpcHRpb24+CiAgIDwvcmRmOlJERj4KPC94OnhtcG1ldGE+Crq03TcAAAGFSURBVDgRhZO9SkNBEIXvTRRNtImFCj5A8KdThCAhNj5I8AG09QHshGAVLaKPoKVdQGuxtBOjjYilGMHo+p3JbFgC4sBhzsycnZ3duzfLsBBCQV4Gr4Mz8AC+HeLK1YeqZA1JW4wvgQ5ILTZIc9KUfLPhxiTK4DpV/cOlLcdpNHbcuQvfBbfgA9w7xJVTrQtkHWsAaShya/toOs4imHSIx7Hb8GiNCRY0R6Nk2ZfzzzzP+0n+hRW5x1GjsKnxH71dH19VFl8E+RiKXquSl1b2KPHAaAhP+CkXxd0UmlGznDRAWtlAnyG4ZgZv5/T4LyfNrBftDfQ8mMM3nNsRnGtK7W5HwG+DCpD1VDzVLNg7eAWbw5rdhd1DEteoP4NoJ2qw7lELf+78Dl8bW3hD7sfreqGyDdNAji0MYQt/AC7BUtJg2ety8Qu0Yj02uaL4BnbAKiiAePMVuGrRLiBW0xGGPwRt4EdRgV+IO8DnPa/RD5P86GdKm6wg2gfTiVBPew+sJTlb8wuNnEcR4dearQAAAABJRU5ErkJggg=="
)

var (
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func main() {
	data := getStatus()
	fmt.Println(data.Status.statusEmoji() + " | image=" + ICON + " | color=" + data.Status.statusColor())
	fmt.Println("---")
	fmt.Println(data.Status.Description + " | href=" + data.Page.Url)
	fmt.Println("Last updated: " + data.Page.lastUpdatedAt())
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
