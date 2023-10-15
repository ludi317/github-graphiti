package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const numRows = 7

var numCols = 53

func main() {

	config, err := loadConfig("config.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commitDate := calendarStartDate()

	message := render(config.Message)

	if config.IsPreview {
		printMessage(message)
		return
	}

	repoURL, err := refreshRemoteDummyRepo(os.Getenv("GH_TOKEN"), config.GitUserName, config.RepoName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	localRepoPath := filepath.Join(config.LocalRepoDir, config.RepoName)

	err = cloneRepo(repoURL, localRepoPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Committing")
	for c := 0; c < numCols; c++ {
		for r := 0; r < numRows; r++ {
			if message[r][c] > 0 {
				err = commit(commitDate, localRepoPath, config)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			commitDate = commitDate.AddDate(0, 0, 1)
		}
		if c%5 == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println()

	fmt.Println("Pushing")
	err = push(localRepoPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Find your message at https://github.com/%s\n", config.GitUserName)
}

func calendarStartDate() time.Time {
	// find the most recent Saturday (bottom right corner of the calendar)
	endDate := time.Now()
	weekDay := endDate.Weekday()
	if weekDay != time.Saturday {
		numCols -= 1
		endDate = endDate.AddDate(0, 0, -int(weekDay+1)%7)
	}
	// work backwards to the first Sunday (top left corner of the calendar)
	startDate := endDate.AddDate(0, 0, -1*numCols*numRows+1)
	return startDate
}
