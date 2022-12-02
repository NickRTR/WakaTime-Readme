package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func allTime(token string) string {
	resp, err := http.Get(fmt.Sprintf("https://wakatime.com/api/v1/users/current/all_time_since_today?api_key=%s", token))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func main() {
	var token string = "naaaa"

	log.Println(allTime(token))
}
