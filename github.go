package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os/exec"
)

func cloneRepo(repoURL, localRepoPath string) error {

	cmd := exec.Command("rm", "-rf", localRepoPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to remove local repo: %v", err)
	}

	cmd = exec.Command("git", "clone", repoURL, localRepoPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to clone repo: %v", err)
	}
	return nil
}

// refreshRemoteDummyRepo deletes and recreates the dummy repository to clear out its contributions from the calendar
func refreshRemoteDummyRepo(token, owner, repoName string) (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// delete the repository
	resp, err := client.Repositories.Delete(ctx, owner, repoName)
	if resp == nil {
		return "", fmt.Errorf("Failed to delete repository %q: %v", repoName, err)
	}
	switch resp.StatusCode {
	case 204:
		fmt.Printf("Repository %q deleted\n", repoName)
	case 404:
		fmt.Printf("Repository %q not found\n", repoName)
	default:
		return "", fmt.Errorf("Failed to delete repository. Status: %s\n", resp.Status)
	}

	// create the repository
	repo := &github.Repository{
		Name:        github.String(repoName),
		Description: github.String("dummy repo for painting github contribution graph"),
	}

	repo, resp, err = client.Repositories.Create(ctx, "", repo)
	if resp == nil {
		return "", fmt.Errorf("Failed to create repository %q: %v", repoName, err)
	}
	switch resp.StatusCode {
	case 201:
		fmt.Printf("Repository %q created\n", repoName)
	default:
		return "", fmt.Errorf("Failed to delete %q repository. Status: %s\n", repoName, resp.Status)
	}
	return *repo.CloneURL, nil
}
