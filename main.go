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

func createGraph(langs languages) {
	empty := "⬜"
	done := "🟩"

	for i := 0; i < 5; i++ {
		l := langs[i]
		percent := math.Round(l.Percent)
		fmt.Printf("%-20s %15s %s %5.2f %%\n", l.Name, l.Text, strings.Repeat(done, int(percent/4))+strings.Repeat(empty, int(25-int(percent/4))), l.Percent)
	}
}

func main() {
	var token string = os.Getenv("WAKATIME_API_KEY")

	fmt.Println("All Time duration: " + allTime(token))

	fmt.Println(strings.Repeat("-", 95))

	languages := last7Days(token)
	createGraph(languages)
}
