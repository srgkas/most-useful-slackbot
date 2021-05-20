package gh

import (
	"context"
	"fmt"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
	"log"
	"strings"
)

// Repo DTO
type Repo struct {
	Owner string
	Name string
}

// Release DTO
type Release struct {
	Repo *Repo
	Tag string
}

// NewRelease create release DTO
// repo: {owner}/{name} srgkas/most-useful-slackbot
// tag: v0.0.1
func NewRelease(repository string, tag string) *Release {
	repo := repoFromString(repository)

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
	targetRelease, err := releaser.getReleaseByTag(r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found target release: %s\n", *targetRelease.TagName)

	err = releaser.uncheckPreRelease(r, targetRelease)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Released version: %s:/%s,%s\n", r.Repo.Owner, r.Repo.Name, r.Tag)

	return nil
}

func NewReleaser(token string) Releaser {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubReleaser{client: client}
}

func repoFromString(repo string) *Repo {
	parts := strings.Split(repo, "/")

	return &Repo{
		Owner: parts[0],
		Name:  parts[1],
	}
}

func (releaser GithubReleaser) uncheckPreRelease(release *Release, ghRelease *github.RepositoryRelease) error {
	ctx := context.Background()

	*ghRelease.Prerelease = false

	ghRelease, _, err := releaser.client.Repositories.EditRelease(
		ctx,
		release.Repo.Owner,
		release.Repo.Name,
		*ghRelease.ID,
		ghRelease,
	)

	return err
}

func (releaser GithubReleaser) getReleaseByTag(r *Release) (*github.RepositoryRelease, error) {
	ctx := context.Background()

	release, _, err := releaser.client.Repositories.GetReleaseByTag(
		ctx,
		r.Repo.Owner,
		r.Repo.Name,
		r.Tag,
	)

	return release, err
}

// Release should be done in go-routines
