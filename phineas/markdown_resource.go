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
	// TODO rebuild w/ root() and base()?

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

func (r *MarkdownResource) base() string {
	return filepath.Base(r.contentPath)
}

func (r *MarkdownResource) root() string {
	root, err := filepath.Rel(r.contentRoot, filepath.Dir(r.contentPath))
	slog.PanicIf(err)
	return root
}

func (r *MarkdownResource) dayArchiveId() string {
	return filepath.Join(r.root(), r.base()[:10])
}

func (r *MarkdownResource) monthArchiveId() string {
	return filepath.Join(r.root(), r.base()[:7])
}

func (r *MarkdownResource) yearArchiveId() string {
	return filepath.Join(r.root(), r.base()[:4])
}
