package phineas

import (
	"log"
	"sort"
)

type Resources []Resource

func (resources Resources) Build() error {
	sort.Sort(resources)

	for _, resource := range resources {
		log.Println("Building", resource.BuildPath())

		if err := resource.Build(); err != nil {
			return err
		}
	}

	return nil
}

func (s Resources) Len() int {
	return len(s)
}

func (s Resources) Less(i, j int) bool {
	return len(s[i].BuildPath()) < len(s[j].BuildPath())
}

func (s Resources) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Index struct {
	buildPath string
	resource  Resource
}

func (r *Index) BuildPath() string {
	return "todo"
}

type Indexes []Index
