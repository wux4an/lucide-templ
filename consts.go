package main

import (
	"encoding/json"
	"net/http"
)

const (
	Version    = "0.555.0"
	GitTreeURL = "https://api.github.com/repos/lucide-icons/lucide/git/trees/" + Version
)

var URL = getUrl()

type TreeResp struct {
	Tree []struct {
		Path string `json:"path"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"tree"`
}

func getUrl() string {
	resp, err := http.Get(GitTreeURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var data TreeResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ""
	}

	for _, item := range data.Tree {
		if item.Path == "icons" && item.Type == "tree" {
			return item.URL
		}
	}

	return ""
}
