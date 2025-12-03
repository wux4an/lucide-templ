package src

import (
	"encoding/json"
	"net/http"
	"strings"
)

type TreeResponse struct {
	Tree []struct {
		Path string `json:"path"`
	} `json:"tree"`
}

func fetchPaths(URL string) []string {
	resp, err := http.Get(URL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var data TreeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	paths := make([]string, 0, len(data.Tree))
	for _, item := range data.Tree {
		paths = append(paths, item.Path)
	}

	return paths
}

func GetSvg(URL string) []string {
	paths := fetchPaths(URL)
	if paths == nil {
		return nil
	}

	svgs := make([]string, 0, len(paths))
	for _, p := range paths {
		if strings.HasSuffix(p, ".svg") {
			svgs = append(svgs, p)
		}
	}
	return svgs
}
