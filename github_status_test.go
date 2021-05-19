package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPageLastUpdatedAt(t *testing.T) {
	page := Page{
		"xxxxx",
		"GitHub",
		"https://www.githubstatus.com",
		"Etc/UTC",
		time.Date(2021, 5, 19, 10, 45, 56, 0, time.UTC),
	}
	result := page.lastUpdatedAt()
	if result != "2021/05/19(Wed) 19:45:56" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}
}

func TestStatusEmoji(t *testing.T) {
	var status Status
	var result string

	status = Status{"none", "All Systems Operational"}
	result = status.statusEmoji()
	if result != "✓" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"minor", "Minor Service Outage"}
	result = status.statusEmoji()
	if result != "!" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"partial", "Partial System Outage"}
	result = status.statusEmoji()
	if result != "!" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"under_maintenance", "Service Under Maintenance"}
	result = status.statusEmoji()
	if result != "!" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"major", "Major Service Outage"}
	result = status.statusEmoji()
	if result != "✗" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"unknown", "Unknown description"}
	result = status.statusEmoji()
	if result != "?" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}
}

func TestStatusColor(t *testing.T) {
	var status Status
	var result string

	status = Status{"none", "All Systems Operational"}
	result = status.statusColor()
	if result != "green" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"minor", "Minor Service Outage"}
	result = status.statusColor()
	if result != "yellow" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"partial", "Partial System Outage"}
	result = status.statusColor()
	if result != "yellow" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"under_maintenance", "Service Under Maintenance"}
	result = status.statusColor()
	if result != "yellow" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"major", "Major Service Outage"}
	result = status.statusColor()
	if result != "red" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}

	status = Status{"unknown", "Unknown description"}
	result = status.statusColor()
	if result != "white" {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", result))
	}
}