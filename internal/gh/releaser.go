package gh

import (
	"context"
	"fmt"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
	"log"
	"strings"
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
	ctx := context.Background()

	owner, repoName := repoToOwnerRepoName(r.Repo)


	// Test API
	releases, _, err := releaser.client.Repositories.ListReleases(ctx, owner, repoName, &github.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(releases)

	fmt.Printf("Released version: %s:%s\n", r.Repo, r.Tag)

	return nil
}

func CreateWithoutClient(token string ) Releaser {
	return &GithubReleaser{client: nil}
}

func CreateReleaser(token string) Releaser {
	// Context might be from outside. TBD
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubReleaser{client: client}
}

func repoToOwnerRepoName(repo string) (string, string) {
	parts := strings.Split(repo, "/")

	return parts[0], parts[1]
}

func getReleaseByTag(releases []*github.RepositoryRelease, tag string) (*github.RepositoryRelease, error) {
	return nil, nil
}

func uncheckPreRelease(release *github.RepositoryRelease) error {
	return nil
}

// Release should be done in go routines