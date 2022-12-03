package main

import (
	"context"
	"fmt"
	"log"
	"strings"

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

func addGraph(client *github.Client, graph string) {
	file, _, _, err := client.Repositories.GetContents(context.Background(), "NickRTR", "WakaTime-Readme", "README.md", nil)
	if err != nil {
		log.Panicln(err)
	}

	readme, err := file.GetContent()
	if err != nil {
		log.Panicln(err)
	}

	editedReadme := strings.Replace(readme, "<!--WakaTime-Start-->\n<!--WakaTime-End-->", fmt.Sprintf("<!--WakaTime-Start-->\n%s<!--WakaTime-End-->", graph), 1)

	b := []byte(editedReadme)
	sha := file.GetSHA()
	message := "Provide README with WakaTime stats"

	updatedFile := github.RepositoryContentFileOptions{
		Message: &message,
		Content: b,
		SHA:     &sha,
	}
	_, _, err = client.Repositories.UpdateFile(context.Background(), "NickRTR", "WakaTime-Readme", "README.md", &updatedFile)
	if err != nil {
		log.Panicln(err)
	}
}