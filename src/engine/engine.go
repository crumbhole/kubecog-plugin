package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

// Engine is a 'class' to hold the values for doing template runs with a single set of variables
// called values, over several golang templated files
type Engine struct {
	Values interface{}
}

// Run will use the Engine's values to templatise one file, in place, given by path
func (e *Engine) Run(path string) error {
	fmt.Printf("Checking path %s\n", path)
	tmpl := template.New("engine")
	filecontents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	tmpl.Parse(string(filecontents))
	outfile, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0400)
	if err != nil {
		return err
	}
	err = tmpl.Execute(outfile, e.Values)
	if err != nil {
		return err
	}

	return nil
}
