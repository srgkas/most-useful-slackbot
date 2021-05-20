package gh

import (
	"github.com/google/go-github/v35/github"
	"log"
)

// Release DTO

type Release struct {
	Repo string
	Tag string
}

func NewRelease(repo string, tag string) *Release {
	return &Release{Repo: repo, Tag: tag}
}

// Releaser code

type Releaser interface {
	Release(r *Release) error
}

type GithubReleaser struct {
	client *github.Client
}

func (releaser GithubReleaser) Release(r *Release) error {
	log.Printf("Released version: %s:%s\n", r.Repo, r.Tag)

	return nil
}

func CreateReleaser(token string) Releaser {
	// Testing
	return &GithubReleaser{client: nil}

	//ctx := context.Background()
	//ts := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: token},
	//)
	//
	//tc := oauth2.NewClient(ctx, ts)
	//
	//client := github.NewClient(tc)
	//
	//return &GithubReleaser{client: client}
}

// Release should be done in go routines