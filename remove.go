package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
)

func removeRepo(c *cli.Context) error {
	if !c.Bool("keep") {
		if c.NArg() == 0 {
			return fmt.Errorf("Please specify a repo to remove")
		}
		folderName := c.Args().First()

		if strings.ContainsAny(folderName, "\\/.") {
			return fmt.Errorf("Repo name cannot contain . / or \\")
		}

		if c.Bool("force") {
			rmerr := os.RemoveAll(folderName)

			if rmerr != nil {
				return fmt.Errorf("Could not remove %v: %v", folderName, rmerr)
			}
		} else {
			repo, err := git.PlainOpen(folderName)
			if err != nil {
				return err
			}
			w, _ := repo.Worktree()
			status, _ := w.Status()
			if !status.IsClean() {
				return fmt.Errorf("Work tree is not clean")
			}

			// // Check for uncommited changes
			// statusCmd := exec.Command("git", "-C", folderName, "status", "--porcelain")
			// output, err := statusCmd.CombinedOutput()

			// if err != nil {
			// 	return fmt.Errorf("Failed to check %v for unstaged changes, err: %v", folderName, err)
			// }

			// if len(output) != 0 {
			// 	return fmt.Errorf("Cannot remove %v: unstaged changes", folderName)
			// }

			// // Check for unapplied stashes
			// stashCommand := exec.Command("git", "-C", folderName, "stash", "list")
			// output, err = stashCommand.CombinedOutput()

			// if err != nil {
			// 	return fmt.Errorf("Failed to check %v for stashes, err: %v", folderName, err)
			// }

			// if len(output) != 0 {
			// 	return fmt.Errorf("Cannot remove %v: unapplied stashes", folderName)
			// }

			// // Check for unpushed commits

			// if err != nil {
			// 	fmt.Printf("Output: %v\nError: %v", string(output), err)
			// } else {
			// 	fmt.Printf("No error, output: %v", output)
			// }
		}
	}

	return nil
}
