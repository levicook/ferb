package phineas

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ArchivePage struct {
	archiveId  string
	buildRoot  string
	pageNumber int
	perPage    int
	resources  Resources
}

func (ap *ArchivePage) Build() error {
	buildPath := ap.BuildPath()

	if err := ensureDirectoryExists(filepath.Dir(buildPath)); err != nil {
		return err
	}

	file, err := os.Create(buildPath)
	defer file.Close()
	return err
}

func (ap *ArchivePage) BuildPath() string {
	path := filepath.Join(
		ap.buildRoot,
		filepath.Dir(ap.archiveId),
		strings.Replace(ap.base(), "-", string(os.PathSeparator), -1),
	)

	if ap.pageNumber > 1 {
		path = filepath.Join(path, "page", fmt.Sprintf("%v", ap.pageNumber))
	}

	return filepath.Join(path, "index.html")
}

func (ap *ArchivePage) base() string {
	return filepath.Base(ap.archiveId)
}
