package src

import (
	"io"
	"net/http"
	"strings"
)

var GetRawSvg = func(rawURL string) string {
	resp, err := http.Get(rawURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

func iconNameFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, ".svg")
	parts := strings.Split(name, "-")

	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

// Inserta los attrs después de <svg
func injectAttrs(svg string) string {
	idx := strings.Index(svg, "<svg")
	if idx == -1 {
		return svg
	}

	insertAt := idx + len("<svg")

	injection := `
    if len(attrs) > 0 {
        { attrs[0]... }
    }`

	return svg[:insertAt] + injection + svg[insertAt:]
}

func MakeTempl(svg string, filename string) string {
	name := iconNameFromFilename(filename)
	svgWithAttrs := injectAttrs(svg)

	return "package icon \n\n" + `templ ` + name + `(attrs ...templ.Attributes) {
` + svgWithAttrs + `
}`
}
