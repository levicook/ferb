package features

import (
	"os"
	"path"

	"github.com/levicook/ferb/phineas"
	. "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
)

func init() {
	var currentApp struct {
		Path string
	}

	Given(`^a successfully built app at "(.+?)"$`, func(p string) {
		currentApp.Path = p

		err := os.Chdir(path.Join("internal/fixtures", currentApp.Path))
		assert.NoError(T, err)

		err = os.RemoveAll("build")
		assert.NoError(T, err)

		err = phineas.Build()
		assert.NoError(T, err)
	})

	When(`^I cd into "(.+?)"$`, func(p string) {
		err := os.Chdir(p)
		assert.NoError(T, err)
	})

	Then(`^the following files should exist:$`, func(table [][]string) {
		for _, row := range table {
			_, err := os.Stat(row[0])
			assert.NoError(T, err)
		}
	})
}
