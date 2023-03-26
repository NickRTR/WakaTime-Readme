package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/NickRTR/WakaTime-Readme/cli"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

func Authenticate(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}

func AddStats(client *github.Client, stats string, user string, repo string) {
	file, _, _, err := client.Repositories.GetContents(context.Background(), user, repo, "README.md", nil)
	if err != nil {
		cli.BrintErr(err.Error())
	}

	readme, err := file.GetContent()
	if err != nil {
		cli.BrintErr(err.Error())
	}

	start := "<!--WakaTime-Start-->"
	stop := "<!--WakaTime-End-->"
	startIndex := strings.Index(readme, start)
	stopIndex := strings.Index(readme, stop) + len(stop)
	sectionBefore := readme[0:startIndex]
	sectionAfter := readme[stopIndex:]

	editedReadme := sectionBefore + fmt.Sprintf("<!--WakaTime-Start-->\n<pre>%s</pre>\n<!--WakaTime-End-->", stats) + sectionAfter

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
		cli.BrintErr(err.Error())
	}

	cli.Brint("Added Graph to README.md")
}
