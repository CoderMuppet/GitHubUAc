package main

import (
	"fmt"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <username>")
		return
	}

	username := os.Args[1]
	GitHubApiUrl := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	events, err := getGitHubEvents(GitHubApiUrl)
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}

	eventPushCount := 0
	// Iterate over the array of events
	for _, event := range events {
		if event.Type == "PushEvent" {
			eventPushCount++

		}
	}
	fmt.Println("Output:")
	fmt.Printf("- Pushed %d commits to %s\n", eventPushCount, events[0].Repo.Name)
}
