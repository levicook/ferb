package phineas

import (
	"log"
	"os"
	"path/filepath"
	"sort"
)

func ensureDirectoryExists(path string) error {
	if err := os.MkdirAll(path, 0777); err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func Build() error {
	var (
		sourceRoot  = "source"
		contentRoot = filepath.Join(sourceRoot, "content")
		buildRoot   = "build"
		//layoutsRoot = filepath.Join(sourceRoot, "layouts")
	)

	ensureDirectoryExists(buildRoot)

	var resources Resources

	err := filepath.Walk(contentRoot, func(contentPath string, contentInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if contentPath == contentRoot {
			return nil
		}

		switch {

		case contentInfo.IsDir():
			resources = append(resources, &DirectoryResource{
				buildRoot:   buildRoot,
				contentRoot: contentRoot,
				contentPath: contentPath,
			})

		case filepath.Ext(contentPath) == ".md":
			resources = append(resources, &MarkdownResource{
				buildRoot:   buildRoot,
				contentRoot: contentRoot,
				contentPath: contentPath,
			})

		case filepath.Ext(contentPath) == ".tmpl":
			resources = append(resources, &TemplateResource{
				buildRoot:   buildRoot,
				contentRoot: contentRoot,
				contentPath: contentPath,
			})

		}

		return nil
	})

	sort.Sort(resources)

	for i := range resources {
		log.Println("Building", resources[i].BuildPath())

		if err = resources[i].Build(); err != nil {
			return err
		}
	}

	return err
}
