package features

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/levicook/ferb/phineas"
	. "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
)

var projectRoot string

func init() {
	var err error
	projectRoot, err = os.Getwd()
	assert.NoError(T, err)
}

func init() {
	var currentApp struct {
		Path string
	}

	var openFile struct {
		Contents string
	}

	Given(`^a successfully built app at "(.+?)"$`, func(p string) {
		currentApp.Path = p

		err := os.Chdir(filepath.Join(projectRoot, "internal/fixtures", currentApp.Path))
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

	And(`^the following files should not exist:$`, func(table [][]string) {
		for _, row := range table {
			_, err := os.Stat(row[0])
			assert.True(T, os.IsNotExist(err))
		}
	})

	When(`^I open "(.+?)"$`, func(filename string) {
		data, err := ioutil.ReadFile(filename)
		assert.NoError(T, err)
		openFile.Contents = string(data)
	})

	Then(`^I should see:$`, func(content string) {
		assert.Contains(T, openFile.Contents, content)
	})
}
