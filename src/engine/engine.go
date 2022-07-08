package engine

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// valuesEnv is the name of the environment variable controlling where to find values yaml
const (
	gomplateDefault = `gomplate`
	gomplateEnv     = `ARGOCD_ENV_KUBECOG_GOMPLATE_PATH`
	urlEnv          = `ARGOCD_ENV_KUBECOG_VALUES_URL_PREFIX`
)

func getEnv(name string, defaultVal string) string {
	result, got := os.LookupEnv(name)
	if !got {
		return defaultVal
	}
	return result
}

// Engine is a 'class' to hold the values for doing template runs with a single set of variables
// called values, over several golang templated files
type Engine struct {
	Values map[string]string
}

// Run will use the Engine's values to templatise one file, in place, given by path
func (e *Engine) Run(path string) error {
	fmt.Printf("Checking path %s\n", path)
	params := make([]string, 0)
	params = append(params, `-f`, path, `-o`, path)
	urlPrefix := getEnv(urlEnv, ``)
	if urlPrefix == `` {
		return fmt.Errorf("Must set %s to URL prefix", urlEnv)
	}
	for name, contextPath := range e.Values {
		params = append(params, `-c`, name+`=`+urlPrefix+contextPath)
	}
	fmt.Printf("Params are %v", params)
	cmd := exec.Command(getEnv(gomplateEnv, gomplateDefault), params...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %v", err, stderr.String())
	}

	return nil
}
