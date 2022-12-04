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
	case "0":
		empty = "⬜"
		done = "🟩"
	case "1":
		empty = "⬜"
		done = "🟨"
	default:
		empty = "⬜"
		done = "🟩"
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
	var token string = os.Getenv("WAKATIME_API_KEY")

	languages := last7Days(token)
	graph := createGraph(languages)

	var GH_TOKEN string = os.Getenv("GH_TOKEN")
	client := authenticate(GH_TOKEN)
	addGraph(client, graph)
}
