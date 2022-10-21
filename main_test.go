package main

import (
	"github.com/crumbhole/kubecog-plugin/src/engine"
	"github.com/crumbhole/kubecog-plugin/src/values"
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
	kubecogRegexp := regexp.MustCompile(`\.kubecog\.yaml$`)
	if kubecogRegexp.MatchString(path) {
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
	cwd, _ := os.Getwd()
	os.Chdir(path)
	os.Setenv(`ARGOCD_ENV_KUBECOG_URL_PREFIX`, `file:///`+cwd+`/`+path+`/`)
	defer os.Chdir(cwd)
	config, err := values.Values()
	if err != nil {
		return err
	}
	s := scanner{engine: engine.Engine{Config: config}}
	err = s.scanDir(`.`)
	if err != nil {
		return err
	}
	return filepath.Walk(`.`, c.checkExpected)
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
