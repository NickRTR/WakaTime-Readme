package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func request(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func last7Days(token string) languages {

	res := request(fmt.Sprintf("https://wakatime.com/api/v1/users/current/stats/last_7_days?api_key=%s&scope=read_stats", token))

	var data stats
	json.Unmarshal([]byte(res), &data)

	return data.Data.Languages
}

func allTime(token string) string {
	res := request(fmt.Sprintf("https://wakatime.com/api/v1/users/current/all_time_since_today?api_key=%s", token))

	var data allTimeData
	json.Unmarshal([]byte(res), &data)

	return data.Data.Text
}

func createGraph(langs languages) string {
	theme := os.Getenv("THEME")

	var empty string
	var done string

	switch theme {
	case "block-green":
		empty = "⬜"
		done = "🟩"
	case "block-yellow":
		empty = "⬜"
		done = "🟨"
	case "block-red":
		empty = "⬜"
		done = "🟥"
	case "block-blue":
		empty = "⬜"
		done = "🟦"
	case "default":
		empty = "░"
		done = "█"
	default:
		empty = "░"
		done = "█"
	}

	var graph string

	for i := 0; i < 5; i++ {
		l := langs[i]
		percent := math.Round(l.Percent)
		graph += fmt.Sprintf("%-15s %15s %s %5.2f %%</br>", l.Name, l.Text, strings.Repeat(done, int(percent/4))+strings.Repeat(empty, int(25-int(percent/4))), l.Percent)
	}

	return graph
}

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

	languages := last7Days(token)
	graph := createGraph(languages)

	var GH_TOKEN string = os.Getenv("GH_TOKEN")
	client := authenticate(GH_TOKEN)
	addGraph(client, graph)
}
