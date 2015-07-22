package phineas

import (
	"os"
	"path/filepath"
	"strings"
)

type Archive struct {
	archiveId string
	buildRoot string
	resources Resources
}

func (a *Archive) Build() error {
	buildPath := a.BuildPath()

	if err := ensureDirectoryExists(filepath.Dir(buildPath)); err != nil {
		return err
	}

	file, err := os.Create(buildPath)
	defer file.Close()
	return err
}

func (a *Archive) BuildPath() string {
	return filepath.Join(
		a.buildRoot,
		filepath.Dir(a.archiveId),
		strings.Replace(a.base(), "-", string(os.PathSeparator), -1),
		"index.html",
	)
}

func (a *Archive) add(r Resource) {
	// TODO
}

func (a *Archive) base() string {
	return filepath.Base(a.archiveId)
}
