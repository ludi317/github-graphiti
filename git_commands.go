package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func commit(date time.Time, localDummyRepo string, config *Config) error {

	dateStr := date.Format("Mon Jan 2 15:04:05 2006")
	var commitStr strings.Builder
	for i := 0; i < config.Intensity; i++ {
		commitStr.WriteString("git commit --allow-empty -m 'Generated commit';")
	}
	cmd := exec.Command("bash", "-c", commitStr.String())
	cmd.Dir = localDummyRepo
	cmd.Env = append(cmd.Env,
		fmt.Sprintf("GIT_AUTHOR_DATE=%s", dateStr),
		fmt.Sprintf("GIT_AUTHOR_EMAIL=%s", config.GitUserEmail),
		fmt.Sprintf("GIT_COMMITTER_DATE=%s", dateStr),
		fmt.Sprintf("GIT_COMMITTER_EMAIL=%s", config.GitUserEmail),
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to commit: %s: %v", out, err)
	}
	return nil
}

func getGitConfigValue(key string) (string, error) {
	cmd := exec.Command("git", "config", "--get", key)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Failed to get git config value of key %s: %v", key, err)
	}
	return strings.TrimSpace(string(output)), nil
}

func push(localRepoPath string) error {
	cmd := exec.Command("git", "push")
	cmd.Dir = localRepoPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to push: %s: %v", out, err)
	}
	return nil
}
