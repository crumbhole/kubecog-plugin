package main

import (
	"github.com/crumbhole/crumblecog-plugin/src/engine"
	"github.com/crumbhole/crumblecog-plugin/src/values"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type scanner struct {
	engine engine.Engine
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
		err := s.engine.Run(path)
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
	vals, err := values.Values()
	if err != nil {
		log.Fatal(err)
	}
	s := scanner{engine: engine.Engine{Values: vals}}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = s.scanDir(dir)
	if err != nil {
		log.Fatal(err)
	}
}