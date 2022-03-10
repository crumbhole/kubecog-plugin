package values

import (
	"os"
	"reflect"
	"testing"
)

func TestDefault(t *testing.T) {
	res, err := Values()
	if err != nil {
		t.Fatal(err)
	}

	expected := make(map[interface{}]interface{})
	expected[`abc`] = `def`
	expected[`foo`] = make(map[interface{}]interface{})
	expected[`foo`].(map[interface{}]interface{})[`bar`] = `123`
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%v not same as expected %v", res, expected)
	}
}

func TestPathEnv(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv(VALUES_ENV, cwd+`/subdir/test.yaml`)
	res, err := Values()
	if err != nil {
		t.Fatal(err)
	}

	expected := make(map[interface{}]interface{})
	expected[`pet`] = make(map[interface{}]interface{})
	expected[`pet`].(map[interface{}]interface{})[`frog`] = `kiss`
	expected[`pet`].(map[interface{}]interface{})[`dog`] = `pat`
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("%v not same as expected %v", res, expected)
	}
}

// func createTestVault(t *testing.T) (net.Listener, *api.Client) {
// 	t.Helper()

// 	// Create an in-memory, unsealed core (the "backend", if you will).
// 	core, keyShares, rootToken := vault.TestCoreUnsealed(t)
// 	_ = keyShares

// 	// Start an HTTP server for the core.
// 	ln, addr := http.TestServer(t, core)

// 	// Create a client that talks to the server, initially authenticating with
// 	// the root token.
// 	conf := api.DefaultConfig()
// 	conf.Address = addr

// 	client, err := api.NewClient(conf)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	client.SetToken(rootToken)

// 	// Setup required secrets, policies, etc.
// 	_, err = client.Logical().Write("secret/data/path", map[string]interface{}{
// 		"data": map[string]interface{}{
// 			"testkey": "
// gunk:
//   goo: no
//   frogs: yes
// ",
// 			"bar": "example",
// 		},
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	return ln, client
// }

// func TestVault(t *testing.T) {
// 	os.Setenv(VALUES_ENV, `secret/data/path~testkey`)
// 	res, err := Values()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := make(map[interface{}]interface{})
// 	expected[`gunk`]=make(map[interface{}]interface{})
// 	expected[`gunk`].(map[interface{}]interface{})[`goo`]=`no`
// 	expected[`gunk`].(map[interface{}]interface{})[`frogs`]=`yes`
// 	if !reflect.DeepEqual(expected, res) {
// 	 	t.Errorf("%v not same as expected %v", res, expected)
// 	}
// }
