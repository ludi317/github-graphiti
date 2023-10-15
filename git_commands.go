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
		commitStr.WriteString("git commit --allow-empty --author '")
		commitStr.WriteString(config.GitUserName)
		commitStr.WriteString(" <")
		commitStr.WriteString(config.GitUserEmail)
		commitStr.WriteString("> ' --date '")
		commitStr.WriteString(dateStr)
		commitStr.WriteString("' -m 'Generated commit'")
		commitStr.WriteString(";")
	}
	cmd := exec.Command("bash", "-c", commitStr.String())
	cmd.Dir = localDummyRepo
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_COMMITTER_DATE=%s", dateStr))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to commit: %v", err)
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
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to push: %v", err)
	}
	return nil
}
