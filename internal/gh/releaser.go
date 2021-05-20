package gh

import (
	"context"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
	"log"
)

type Releaser interface {
	Release(repo string, tag string) error
}

type GithubReleaser struct {
	client *github.Client
}

func (r GithubReleaser) Release(repo string, tag string) error {
	log.Printf("Released version: %s:%s\n", repo, tag)

	return nil
}

func CreateReleaser(token string) Releaser {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubReleaser{client: client}
}