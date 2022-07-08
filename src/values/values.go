package values

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"syscall"
)

// kubecogEnable is the file that enables kubecog and
// tells the system which cog files to read from
const kubecogEnable = `./.kubecog.yaml`

type Delims struct {
	Left  string `json:"left,omitempty"`
	Right string `json:"right,omitempty"`
}

type Kubecog struct {
	Kubecog    map[string]string `json:"kubecog"`
	Delimiters *Delims           `json:"delimiters,omitempty"`
}

// Values is a function to get you a map of name: values.yaml
func Values() (*Kubecog, error) {
	contents, err := ioutil.ReadFile(kubecogEnable)
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		print("No .kubecog.yaml\n")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var kubecog Kubecog
	err = yaml.Unmarshal(contents, &kubecog)
	return &kubecog, err
}
