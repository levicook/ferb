package phineas

import (
	"path/filepath"

	"github.com/levicook/slog"
)

type DirectoryResource struct {
	buildRoot   string
	contentRoot string
	contentPath string
}

func (r *DirectoryResource) Build() error {
	return ensureDirectoryExists(r.BuildPath())
}

func (r *DirectoryResource) BuildPath() string {
	commonPath, err := filepath.Rel(r.contentRoot, r.contentPath)
	slog.PanicIf(err)

	return filepath.Join(r.buildRoot, commonPath)
}
