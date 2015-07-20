package phineas

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/levicook/slog"
)

type TemplateResource struct {
	buildRoot   string
	contentRoot string
	contentPath string
}

func (r *TemplateResource) Build() error {
	buildPath := r.BuildPath()

	if err := ensureDirectoryExists(filepath.Dir(buildPath)); err != nil {
		return err
	}

	file, err := os.Create(buildPath)
	defer file.Close()
	return err
}

func (r *TemplateResource) BuildPath() string {
	path, err := filepath.Rel(r.contentRoot, r.contentPath)
	slog.PanicIf(err)

	path = strings.TrimSuffix(path, ".tmpl")
	base := filepath.Base(path)

	if filepath.Ext(base) == ".html" && base != "index.html" {
		path = strings.TrimSuffix(path, ".html")
		path = filepath.Join(path, "index.html")
	}

	return filepath.Join(r.buildRoot, path)
}
