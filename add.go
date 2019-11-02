package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/urfave/cli"
)

var nameRegex = regexp.MustCompile("/([a-zA-Z0-9_\\-]+)(\\.git)?(/)?$")

func getRepoName(url string) string {
	matches := nameRegex.FindStringSubmatch(url)
	fmt.Println(url, matches)
	repoName := matches[1]
	return repoName
}

func cloneRepo() {

}

func addRepo(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("Please specify a repo URL")
	}
	repoURL := c.Args().First()
	folderName := getRepoName(repoURL)

	if c.String("dir") != "" {
		folderName = c.String("dir")
	}

	if stat, err := os.Stat(folderName); err == nil && stat.IsDir() {
		return fmt.Errorf("%v is already a direcotry", folderName)
	}

	fmt.Printf("Cloning %v into %v\n", repoURL, folderName)

	cloneCmd := exec.Command("git", "clone", repoURL, folderName)
	output, err := cloneCmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))

		if stat, err := os.Stat(folderName); err == nil && stat.IsDir() {
			os.RemoveAll(folderName)
		}

		return fmt.Errorf("Failed to clone %v: %v", repoURL, err)
	}

	return nil
}
