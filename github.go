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

func addGraph(client *github.Client, graph string, user string, repo string) {
	file, _, _, err := client.Repositories.GetContents(context.Background(), user, repo, "README.md", nil)
	if err != nil {
		log.Panicln(err)
	}

	readme, err := file.GetContent()
	if err != nil {
		log.Panicln(err)
	}

	start := "<!--WakaTime-Start-->"
	stop := "<!--WakaTime-End-->"
	startIndex := strings.Index(readme, start)
	stopIndex := strings.Index(readme, stop) + len(stop)
	sectionBefore := readme[0:startIndex]
	sectionAfter := readme[stopIndex:]

	editedReadme := sectionBefore + fmt.Sprintf("<!--WakaTime-Start-->\n<pre>%s</pre>\n<!--WakaTime-End-->", graph) + sectionAfter

	b := []byte(editedReadme)
	sha := file.GetSHA()
	message := "Provide README with WakaTime statistics"

	updatedFile := github.RepositoryContentFileOptions{
		Message: &message,
		Content: b,
		SHA:     &sha,
	}

	_, _, err = client.Repositories.UpdateFile(context.Background(), user, repo, "README.md", &updatedFile)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Added Graph to README.md")
}
