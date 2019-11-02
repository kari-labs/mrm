package main

type repo struct {
	URL       string  `json:"url"`
	Directory string  `json:"dir"`
	Branch    *string `json:"branch,omitempty"`
}

type config struct {
	Repos repo `json:"repos"`
}

func addRepoToConfig(repo repo) {

}

func saveConfig(config *config) {

}
