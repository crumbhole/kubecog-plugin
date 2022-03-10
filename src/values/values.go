package values

import (
	"github.com/crumbhole/argocd-vault-replacer/src/bwvaluesource"
	"github.com/crumbhole/argocd-vault-replacer/src/substitution"
	"github.com/crumbhole/argocd-vault-replacer/src/vaultvaluesource"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
)

const VALUES_DEFAULT_PATH = `./cogvalues.yaml`
const VALUES_ENV = `COG_VALUES_PATH`

func getPath() string {
	if envpath, pathpresent := os.LookupEnv(VALUES_ENV); pathpresent {
		return envpath
	} else {
		return VALUES_DEFAULT_PATH
	}
}

func tryLocalFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func tryRemote(path string, key string) ([]byte, error) {
	var vs substitution.ValueSource
	if _, bwpresent := os.LookupEnv(`BW_SESSION`); bwpresent {
		vs = bwvaluesource.BitwardenValueSource{}
	} else {
		vs = vaultvaluesource.VaultValueSource{}
	}
	val, err := vs.GetValue([]byte(path), []byte(key))
	return *val, err
}

func Values() (interface{}, error) {
	path := getPath()
	reSplit := regexp.MustCompile(`\s*\~\s*`)
	splitPath := reSplit.Split(string(path), 2)
	var filecontents []byte
	var err error
	if len(splitPath) == 2 {
		filecontents, err = tryRemote(splitPath[0], splitPath[1])
		if err != nil {
			return nil, err
		}
	} else {
		filecontents, err = tryLocalFile(path)
		if err != nil {
			return nil, err
		}
	}
	var values interface{}
	err = yaml.Unmarshal(filecontents, &values)

	return values, err
}
