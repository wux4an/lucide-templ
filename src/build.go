package src

import (
	"os"
	"path/filepath"
	"strings"
)

func Build(templ string, dir string, filename string) {
	// Crear directorio si no existe

	// .svg → .templ
	base := strings.TrimSuffix(filename, ".svg")
	filename = base + ".templ"

	path := filepath.Join(dir, filename)

	// Escribir archivo
	_ = os.WriteFile(path, []byte(templ), 0o644)
}
