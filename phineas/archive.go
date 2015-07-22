package phineas

type Archive struct {
	archiveId string
	buildRoot string
	resources Resources
}

func (a *Archive) add(r Resource) {
	a.resources = append(a.resources, r)
}

func (a *Archive) ArchivePages() (pages []ArchivePage) {
	// TODO perPage from config / frontmatter?
	// TODO sort resources
	pageNumber := 1
	perPage := 1

	page := ArchivePage{
		archiveId:  a.archiveId,
		buildRoot:  a.buildRoot,
		perPage:    perPage,
		pageNumber: pageNumber,
	}

	for _, resource := range a.resources {
		page.resources = append(page.resources, resource)

		if len(page.resources) == perPage {
			pageNumber++
			pages = append(pages, page)

			page = ArchivePage{
				archiveId:  a.archiveId,
				buildRoot:  a.buildRoot,
				perPage:    perPage,
				pageNumber: pageNumber,
			}
		}
	}

	return pages
}
