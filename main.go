package main

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/wux4an/lucide-templ/src"
)

var RawURL = func(version string, file string) string {
	return "https://raw.githubusercontent.com/lucide-icons/lucide/" + version + "/icons/" + file
}

var (
	wg      sync.WaitGroup
	workers = runtime.NumCPU() * 4
)

func main() {
	buildDir := "./build/icons"
	os.MkdirAll(buildDir, 0o755)
	content := "module github.com/wux4an/lucide-templ\n\ngo 1.24.3\n\nrequire github.com/a-h/templ v0.3.960\n"
	path := filepath.Join("./build/", "go.mod")

	// Generate templs
	os.WriteFile(path, []byte(content), 0o644)

	svgs := src.GetSvg(URL)
	jobs := make(chan string, len(svgs))

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for s := range jobs {
				rawSvg := src.GetRawSvg(RawURL(Version, s))
				svg := src.MakeTempl(rawSvg, s)
				src.Build(svg, buildDir, s)
			}
		}()
	}

	for _, s := range svgs {
		jobs <- s
	}
	close(jobs)

	wg.Wait()
}
