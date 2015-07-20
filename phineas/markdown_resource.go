package phineas

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/levicook/slog"
)

type MarkdownResource struct {
	buildRoot   string
	contentRoot string
	contentPath string
}

func (r *MarkdownResource) Build() error {
	buildPath := r.BuildPath()

	if err := ensureDirectoryExists(filepath.Dir(buildPath)); err != nil {
		return err
	}

	file, err := os.Create(buildPath)
	defer file.Close()
	return err
}

func (r *MarkdownResource) BuildPath() string {
	path, err := filepath.Rel(r.contentRoot, r.contentPath)
	slog.PanicIf(err)

	base := strings.TrimSuffix(filepath.Base(path), ".md")
	base = strings.Replace(base, "-", string(os.PathSeparator), 3)

	return filepath.Join(
		r.buildRoot,
		filepath.Dir(path),
		base,
		fmt.Sprintf("index.html"),
	)
}
