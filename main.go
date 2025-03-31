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

	repoName := events[0].Repo.Name
	eventPushCount := 0
	fmt.Println("Output:")
	// Iterate over the array of events
	for i, event := range events {
		if events[i].Repo.Name == repoName {
			if event.Type == "PushEvent" {
				eventPushCount++
			}
		} else {
			if eventPushCount > 0 {
				fmt.Printf("- Pushed %d commits to %s\n", eventPushCount, events[i-1].Repo.Name)
			}
			eventPushCount = 0
			repoName = events[i].Repo.Name
		}

	}

}
