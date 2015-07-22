package phineas

import (
	"os"
	"path/filepath"
)

func ensureDirectoryExists(path string) error {
	if err := os.MkdirAll(path, 0777); err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func Build() error {
	var (
		//layoutsRoot = filepath.Join(sourceRoot, "layouts")

		sourceRoot  = "source"
		contentRoot = filepath.Join(sourceRoot, "content")
		buildRoot   = "build"
	)

	ensureDirectoryExists(buildRoot)

	var (
		resources Resources
		archives  = make(map[string]*Archive)
	)

	err := filepath.Walk(contentRoot, func(contentPath string, contentInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if contentPath == contentRoot {
			return nil
		}

		joinArchive := func(archiveId string, resource Resource) {
			if _, ok := archives[archiveId]; !ok {
				archives[archiveId] = &Archive{
					archiveId: archiveId,
					buildRoot: buildRoot,
				}
			}

			archives[archiveId].add(resource)
		}

		switch {

		case contentInfo.IsDir():
			resources = append(resources, &DirectoryResource{buildRoot: buildRoot, contentRoot: contentRoot, contentPath: contentPath})

		case filepath.Ext(contentPath) == ".md":
			resource := &MarkdownResource{buildRoot: buildRoot, contentRoot: contentRoot, contentPath: contentPath}
			joinArchive(resource.dayArchiveId(), resource)
			joinArchive(resource.monthArchiveId(), resource)
			joinArchive(resource.yearArchiveId(), resource)
			resources = append(resources, resource)

		case filepath.Ext(contentPath) == ".tmpl":
			resource := &TemplateResource{buildRoot: buildRoot, contentRoot: contentRoot, contentPath: contentPath}
			resources = append(resources, resource)

		}

		return nil
	})

	if err != nil {
		return err
	}

	for i := range archives {

		archivePages := archives[i].ArchivePages()

		for j := range archivePages {
			resources = append(resources, &archivePages[j])
		}
	}

	return resources.Build()
}
