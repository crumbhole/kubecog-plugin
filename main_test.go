package main

import (
	"github.com/crumbhole/crumblecog-plugin/src/engine"
	"github.com/crumbhole/crumblecog-plugin/src/values"
	//	"bytes"
	//  "fmt"
	"github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

const (
	testsPath     = "test/"
	testsPathCopy = "test_copy/"
)

type checker struct {
	t *testing.T
}

func (c *checker) checkExpected(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	fileRegexp := regexp.MustCompile(`\.yaml$`)
	if fileRegexp.MatchString(path) {
		expectedPath := strings.Replace(path, `.yaml`, `.expectyaml`, 1)
		got, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		expected, err := ioutil.ReadFile(expectedPath)
		if err != nil {
			return err
		}
		if string(got) != string(expected) {
			c.t.Errorf("Got \n%s\n and expected \n%s\n for %s\n",
				got,
				expected,
				path)
		}
	}
	return nil
}

func (c *checker) checkDir(path string) error {
	os.Setenv(`COG_VALUES_PATH`, path+`/cogvalues.yml`)
	vals, err := values.Values()
	if err != nil {
		return err
	}
	s := scanner{engine: engine.Engine{Values: vals}}
	err = s.scanDir(path)
	if err != nil {
		return err
	}
	return filepath.Walk(path, c.checkExpected)
}

// Finds directories under ./test and evaluates all the .yaml/.ymls
func TestDirectories(t *testing.T) {
	opt := copy.Options{
		OnDirExists: func(_ string, _ string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err := os.RemoveAll(testsPathCopy)
	if err != nil {
		t.Error(err)
	}
	err = copy.Copy(testsPath, testsPathCopy, opt)
	if err != nil {
		t.Error(err)
	}
	dirs, err := ioutil.ReadDir(testsPathCopy)
	if err != nil {
		t.Error(err)
	}

	checker := checker{t: t}
	for _, d := range dirs {
		if d.IsDir() {
			t.Run(d.Name(), func(t *testing.T) {
				err := checker.checkDir(testsPathCopy + d.Name())
				if err != nil {
					t.Error(err)
				}
			})
		}
	}
}