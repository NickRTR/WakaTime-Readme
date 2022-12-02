package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

func authenticate(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}

func addGraph(client *github.Client, graph [5]string) {
	file, _, _, err := client.Repositories.GetContents(context.Background(), "NickRTR", "WakaTime-Readme", "README.md", nil)
	if err != nil {
		log.Panicln(err)
	}

	readme, err := file.GetContent()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(readme)

	//client.Git.CreateCommit(context.Background(), nil, "WakaTime-Readme", commit)
}
