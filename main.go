package main

import (
	"log"
	"os"
	"strings"

	"github.com/NickRTR/WakaTime-Readme/github"
	"github.com/NickRTR/WakaTime-Readme/stats/requests"
	"github.com/joho/godotenv"
)

func main() {
	// environment variables for local development
	if len(os.Args) > 1 {
		if os.Args[1] == "test" {
			error := godotenv.Load(".env")
			if error != nil {
				log.Fatalln("Could not load .env file")
			}
		}
	}

	var token string = os.Getenv("WAKATIME_API_KEY")
	var GH_TOKEN string = os.Getenv("GH_TOKEN")
	target := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	user := target[0]
	repo := target[1]

	languages := requests.Last7Days(token)
	graph := github.Create7DaysGraph(languages)

	allTime := requests.AllTime(token)
	allTimeMarkdown := github.CreateAllTimeData(allTime)

	markdown := graph + allTimeMarkdown

	client := github.Authenticate(GH_TOKEN)
	github.AddStats(client, markdown, user, repo)
}
