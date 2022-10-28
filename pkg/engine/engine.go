package engine

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/crumbhole/kubecog-plugin/pkg/kubecogConfig"
)

// valuesEnv is the name of the environment variable controlling where to find values yaml
const (
	gomplateDefault = `gomplate`
	gomplateEnv     = `ARGOCD_ENV_KUBECOG_GOMPLATE_PATH`
	urlEnv          = `ARGOCD_ENV_KUBECOG_URL_PREFIX`
	leftDefault     = `[[`
	rightDefault    = `]]`
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
	Config *kubecogConfig.Kubecog
}

// Run will use the Engine's values to templatise one file, in place, given by path
func (e *Engine) Run(path string) error {
	kubecogRegexp := regexp.MustCompile(`\.kubecog\.yaml$`)
	if kubecogRegexp.MatchString(path) {
		return nil
	}
	params := make([]string, 0)
	leftDelim := leftDefault
	rightDelim := rightDefault
	if e.Config.Delimiters != nil && e.Config.Delimiters.Left != `` {
		leftDelim = e.Config.Delimiters.Left
	}
	if e.Config.Delimiters != nil && e.Config.Delimiters.Right != `` {
		rightDelim = e.Config.Delimiters.Right
	}
	params = append(params, `-f`, path, `-o`, path, `--left-delim`, leftDelim, `--right-delim`, rightDelim)
	urlPrefix := getEnv(urlEnv, ``)
	if urlPrefix == `` {
		return fmt.Errorf("Must set %s to URL prefix", urlEnv)
	}
	for name, contextPath := range e.Config.Kubecog {
		params = append(params, `-c`, name+`=`+urlPrefix+contextPath)
	}
	cmd := exec.Command(getEnv(gomplateEnv, gomplateDefault), params...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %v", err, stderr.String())
	}

	return nil
}
