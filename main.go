package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func main() {
	var token string = ""

	log.Println("All Time duration: " + allTime(token))
	languages := last7Days(token)
	for _, l := range languages {
		fmt.Println(l.Name)
	}
}
