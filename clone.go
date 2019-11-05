package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func cloneCommand(c *cli.Context) error {
	configPath := c.String("config")
	if configPath == "" {
		configPath = "mrm.conf"
	}

	config, err := loadConfig(configPath)
	defer config.close()
	if err != nil {
		return fmt.Errorf("Failed to load config \"%v\", err: %v", configPath, err)
	}

	for _, repo := range config.Repos {
		if stat, err := os.Stat(repo.Directory); err == nil && stat.IsDir() {
			fmt.Printf("%v is already a direcotry, skipping\n", repo.Directory)
			continue
		}
		fmt.Printf("Cloning %v into %v\n", repo.URL, repo.Directory)

		err := cloneRepo(repo.URL, repo.Directory)
		if err != nil {
			os.RemoveAll(repo.Directory)
			fmt.Printf("Failed to clone %v into %v\n", repo.URL, repo.Directory)
			fmt.Println(err)
		}

		fmt.Printf("Successfully cloned %v\n", repo.Directory)
	}

	return err
}
