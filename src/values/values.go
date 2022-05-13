package values

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"syscall"
)

// valuesEnv is the name of the environment variable controlling where to find values yaml
const valuesEnv = `COG_VALUES_PATH`

// kubecogEnable is the file that enables kubecog and
// tells the system which cog files to read from
const kubecogEnable = `./.kubecog.yaml`

func getBasePath() string {
	if envpath, pathpresent := os.LookupEnv(valuesEnv); pathpresent {
		return envpath
	}
	cwd, _ := os.Getwd()
	return cwd + `/`
}

type kubecog struct {
	Kubecog []string `json:"kubecog"`
}

func valuePaths() ([]string, error) {
	contents, err := ioutil.ReadFile(kubecogEnable)
	if err != nil {
		return nil, err
	}
	var kubecogPaths kubecog
	err = yaml.Unmarshal(contents, &kubecogPaths)
	return kubecogPaths.Kubecog, err
}

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

// Values is a function to get you a set of values returned via the interface{} which
// have been extracted from one or more .yaml files coming from a local file
func Values() (map[string]interface{}, error) {
	paths, err := valuePaths()
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		print("No .kubecog.yaml\n")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	basePath, err := fromGit()
	if err != nil {
		return nil, err
	}
	if basePath == `` {
		basePath = getBasePath()
	}
	var values map[string]interface{}
	for _, path := range paths {
		readPath := basePath + path
		fmt.Printf("|Reading %v\n", readPath)
		contents, err := ioutil.ReadFile(readPath)
		if err != nil {
			return nil, err
		}
		fmt.Printf("|Contents %v\n", contents)
		var theseValues map[string]interface{}
		err = yaml.Unmarshal(contents, &theseValues)
		if err != nil {
			return nil, err
		}
		fmt.Printf("|Got %v\n", theseValues)
		values = mergeMaps(values, theseValues)
		fmt.Printf("|Finally %v\n", values)
	}

	return values, err
}
