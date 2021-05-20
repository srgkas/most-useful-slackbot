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
	RepoOwner string
	RepoName string
	Tag string
}

func NewRelease(repo string, tag string) *Release {
	owner, repoName := repoToOwnerRepoName(repo)

	return &Release{RepoOwner: owner, RepoName: repoName, Tag: tag}
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

	// TODO: load all releases in separate go-routine because all releases are paginated
	releases, _, err := releaser.client.Repositories.ListReleases(
		ctx,
		r.RepoOwner,
		r.RepoName,
		&github.ListOptions{},
	)

	if err != nil {
		log.Fatal(err)
	}

	targetRelease, err := getReleaseByTag(releases, r.Tag)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found target release: %s\n", *targetRelease.TagName)

	err = releaser.uncheckPreRelease(r, targetRelease)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Released version: %s:/%s,%s\n", r.RepoOwner, r.RepoName, r.Tag)

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
	for _, release := range releases {
		if *release.TagName == tag {
			return release, nil
		}
	}

	return nil, fmt.Errorf("no release found with tag: %s", tag)
}

func (releaser GithubReleaser) uncheckPreRelease(release *Release, ghRelease *github.RepositoryRelease) error {
	ctx := context.Background()

	*ghRelease.Prerelease = false

	ghRelease, _, err := releaser.client.Repositories.EditRelease(
		ctx,
		release.RepoOwner,
		release.RepoName,
		*ghRelease.ID,
		ghRelease,
	)

	return err
}

func (releaser GithubReleaser) loadReleases(r *Release) ([]*github.RepositoryRelease, error) {
	//TODO: implement
	return nil, nil
}

// Release should be done in go routines