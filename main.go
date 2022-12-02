package main

import (
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

func last7Days(token string) string {
	return request(fmt.Sprintf("https://wakatime.com/api/v1/users/current/stats/last_7_days?api_key=%s", token))
}

func allTime(token string) string {
	return request(fmt.Sprintf("https://wakatime.com/api/v1/users/current/all_time_since_today?api_key=%s", token))
}

func main() {
	var token string = "naaaa"

	log.Println(allTime(token))
	log.Println(last7Days(token))
}
