package kubecogConfig

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func chdirRelative(relPath string) error {
	_, callerFile, _, _ := runtime.Caller(0)
	callerPath := filepath.Dir(callerFile)
	return os.Chdir(filepath.Join(callerPath, relPath))
}

func TestNoFile(t *testing.T) {
	err := chdirRelative("test/noFile")
	if err != nil {
		t.Error(err)
	}
	res, err := Values()
	if res != nil {
		t.Error("No .kubecog.yaml didn't result in empty result")
	}
	if err != nil {
		t.Error("No .kubecog.yaml didn't result in empty error")
	}
}

func TestErrors(t *testing.T) {
	tests := map[string]error{
		`test/wrongVersion`: KubecogUnknownVersion,
	}
	for testDir, error := range tests {
		t.Run(testDir, func(t *testing.T) {
			err := chdirRelative(testDir)
			if err != nil {
				t.Error(err)
			}
			res, err := Values()
			if res != nil {
				t.Error("Erroring .kubecog.yaml didn't result in empty result")
			}
			if err != error {
				t.Errorf("Error <%s> in .kubecog.yaml didn't match expected <%s>", err, error)
			}
		})
	}
}

func TestVersions(t *testing.T) {
	tests := map[string]Kubecog{
		`test/v1alpha1`: {
			APIVersion: v1alpha1,
			Kubecog: map[string]string{
				`foo`: `bar.yaml`,
			},
			Delimiters: &Delims{
				Left: `x`,
			},
		},
	}
	for testDir, expected := range tests {
		t.Run(testDir, func(t *testing.T) {
			err := chdirRelative(`test/v1alpha1`)
			if err != nil {
				t.Error(err)
			}
			res, err := Values()
			if err != nil {
				t.Errorf("Unexpected error reading .kubecog.yaml <%s>", err)
			} else if res == nil {
				t.Error("Unexpectedly empty result reading .kubecog.yaml")
			} else if !cmp.Equal(*res, expected) {
				t.Errorf("Result <%v> doesn't match expected <%v>", *res, expected)
			}
		})
	}
}

// func TestFailures(t *testing.T) {
// 	testt = map[string]error
// 	opt := copy.Options{
// 		OnDirExists: func(_ string, _ string) copy.DirExistsAction {
// 			return copy.Replace
// 		},
// 	}
// 	err := os.RemoveAll(testsPathCopy)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = copy.Copy(testsPath, testsPathCopy, opt)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	dirs, err := ioutil.ReadDir(testsPathCopy)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	checker := checker{t: t}
// 	for _, d := range dirs {
// 		if d.IsDir() {
// 			t.Run(d.Name(), func(t *testing.T) {
// 				err := checker.checkDir(testsPathCopy + d.Name())
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			})
// 		}
// 	}
// }
