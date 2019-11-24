package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
)

func removeRepo(c *cli.Context) error {
	configPath := c.String("config")
	if configPath == "" {
		configPath = "mrm.conf"
	}

	config, err := loadConfig(configPath)
	defer config.close()
	if err != nil {
		return fmt.Errorf("Failed to load config \"%v\", err: %v", configPath, err)
	}

	if c.NArg() == 0 {
		return fmt.Errorf("Please specify a repo to remove")
	}
	folderName := c.Args().First()

	err = config.removeRepo(folderName)
	if err != nil {
		return fmt.Errorf("Could not remove %v, err: %v", folderName, err)
	}
	fmt.Printf("Removed %v from config\n", folderName)

	if !c.Bool("keep") {
		if strings.ContainsAny(folderName, "\\/.") {
			return fmt.Errorf("Repo name cannot contain . / or \\")
		}

		if c.Bool("force") {
			rmerr := os.RemoveAll(folderName)

			if rmerr != nil {
				return fmt.Errorf("Could not delete %v: %v", folderName, rmerr)
			}
		} else {
			repo, err := git.PlainOpen(folderName)
			if err != nil {
				return err
			}

			w, wtErr := repo.Worktree()
			if wtErr != nil {
				return wtErr
			}

			status, stErr := w.Status()
			if stErr != nil {
				return stErr
			}

			if !status.IsClean() {
				return fmt.Errorf("There are current unstaged changes on %v", folderName)
			}

			fmt.Printf("MRM cannot automatically check the repo for unpushed changes or stashes, please verify that deleting this repo (%v) is the intended action or add --keep to remove the repo from the config but not delte the directory\n", folderName)
			fmt.Println("Type \"yes\" to confirm deletion, or anything else to cancel")

			s := bufio.NewReader(os.Stdin)
			resp, err := s.ReadString('\n')

			if err != nil {
				return err
			}

			if strings.Trim(resp, "\r\n") != "yes" {
				return fmt.Errorf("Action cancelled by user")
			}

			rmerr := os.RemoveAll(folderName)

			if rmerr != nil {
				return fmt.Errorf("Could not delete %v: %v", folderName, rmerr)
			}

			fmt.Printf("Deleted %v\n", folderName)
		}
	}

	return nil
}
