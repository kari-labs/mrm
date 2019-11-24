package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type repo struct {
	URL       string  `json:"url"`
	Directory string  `json:"dir"`
	Branch    *string `json:"branch,omitempty"`
}

type config struct {
	file  *os.File `json:"-"`
	Repos []repo   `json:"repos"`
}

var globalConfig config

func (cfg *config) addRepo(repo repo) error {
	cfg.Repos = append(cfg.Repos, repo)
	err := cfg.save()
	return err
}

func (cfg *config) removeRepo(dirName string) error {
	for i, repo := range cfg.Repos {
		if repo.Directory == dirName {
			cfg.Repos = append(cfg.Repos[:i], cfg.Repos[i+1:]...)
		}
	}
	err := cfg.save()
	return err
}

func (cfg *config) close() {
	cfg.file.Close()
}

func (cfg *config) save() error {
	jsonData, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	err = cfg.file.Truncate(0)
	if err != nil {
		return err
	}
	cfg.file.Seek(0, 0)

	_, err = cfg.file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func loadConfig(configPath string) (*config, error) {
	cfgFile, err := os.OpenFile(configPath, os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		return nil, err
	}

	var cfg config
	if len(data) != 0 {
		err = json.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
	}

	cfg.file = cfgFile

	return &cfg, nil
}
