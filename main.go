package main

import (
	"log"
	"os"
	"strings"

	"github.com/NickRTR/WakaTime-Readme/github"
	"github.com/NickRTR/WakaTime-Readme/stats/requests"
	"github.com/joho/godotenv"
)

func init() {
	// environment variables for local development
	if len(os.Args) > 1 {
		if os.Args[1] == "test" {
			error := godotenv.Load(".env")
			if error != nil {
				log.Fatalln("Could not load .env file")
			}
		}
	}
}

func main() {
	var token string = os.Getenv("WAKATIME_API_KEY")
	var GH_TOKEN string = os.Getenv("GH_TOKEN")
	target := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	user := target[0]
	repo := target[1]

	// request stats of the last 7 days
	languages := requests.Last7Days(token)
	sevenDaysStats := github.Format7DaysStats(languages)

	// request all time stats
	allTime := requests.AllTime(token)
	allTimeStats := github.FormatAllTimeStats(allTime)

	markdown := sevenDaysStats + allTimeStats

	client := github.Authenticate(GH_TOKEN)
	github.AddStats(client, markdown, user, repo)
}
