package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	RepoName     string
	LocalRepoDir string
	Message      string
	Intensity    int
	GitUserName  string
	GitUserEmail string
	IsPreview    bool
}

func loadConfig(filename string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return &Config{}, fmt.Errorf("Error decoding config file: %v", err)
	}

	username, err := getGitConfigValue("user.name")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.GitUserName = username

	useremail, err := getGitConfigValue("user.email")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.GitUserEmail = useremail

	return &config, nil
}
