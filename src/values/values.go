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

type kubecog struct {
	Kubecog map[string]string `json:"kubecog"`
}

// Values is a function to get you a map of name: values.yaml
func Values() (map[string]string, error) {
	contents, err := ioutil.ReadFile(kubecogEnable)
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		print("No .kubecog.yaml\n")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var kubecogPaths kubecog
	err = yaml.Unmarshal(contents, &kubecogPaths)
	return kubecogPaths.Kubecog, err
}
