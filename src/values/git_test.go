package values

import (
	"os"
	"testing"
)

func TestGit(t *testing.T) {
	os.Setenv(gitURLEnv, `https://github.com/crumbhole/kubecog-example`)
	dir, err := fromGit()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Directory %s\n", dir)
	_, err = os.Stat(dir + `/cogvalues.yaml`)
	if err != nil {
		t.Fatal(err)
	}
}
