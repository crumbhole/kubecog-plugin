package values

import (
	"os"
	"reflect"
	"testing"
)

func TestDefaultValues(t *testing.T) {
	os.Unsetenv(gitURLEnv)
	res, err := Values()
	if err != nil {
		t.Fatal(err)
	}

	expected := make(map[string]interface{})
	expected[`abc`] = `def`
	expected[`foo`] = make(map[interface{}]interface{})
	expected[`foo`].(map[interface{}]interface{})[`bar`] = `123`
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%v not same as expected %v", res, expected)
	}
}

func TestPathEnv(t *testing.T) {
	os.Unsetenv(gitURLEnv)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv(valuesEnv, cwd+`/subdir/`)
	res, err := Values()
	if err != nil {
		t.Fatal(err)
	}

	expected := make(map[string]interface{})
	expected[`pet`] = make(map[interface{}]interface{})
	expected[`pet`].(map[interface{}]interface{})[`frog`] = `kiss`
	expected[`pet`].(map[interface{}]interface{})[`dog`] = `pat`
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%v not same as expected %v", res, expected)
	}
}
