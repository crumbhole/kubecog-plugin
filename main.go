package main

import (
	"bufio"
	"fmt"
	"github.com/crumbhole/crumblecog-plugin/src/substitution"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type scanner struct{}

func (s *scanner) process(input []byte) error {
	subst := substitution.Substitutor{}
	modifiedcontents, err := subst.Substitute(input)
	if err != nil {
		return err
	}
	fmt.Printf("---\n%s\n", modifiedcontents)
	return nil
}

func (s *scanner) scanFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	fileRegexp := regexp.MustCompile(`\.ya?ml$`)
	if fileRegexp.MatchString(path) {
		filecontents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		err = s.process(filecontents)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *scanner) scanDir(path string) error {
	return filepath.Walk(path, s.scanFile)
}

func main() {
	stat, _ := os.Stdin.Stat()
	s := scanner{}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		filecontents, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		err = s.process(filecontents)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		err = s.scanDir(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
}
