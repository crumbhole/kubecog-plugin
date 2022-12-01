package kubecogconfig

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"syscall"
)

// kubecogEnable is the file that enables kubecog and
// tells the system which cog files to read from
const kubecogEnable = `./.kubecog.yaml`

// ErrKubecogUnknownVersion is the error returned when .kubecog.yaml doesn't know the Version
var ErrKubecogUnknownVersion = errors.New(".kubecog.yaml does not have a known Version")

// Delims is the type for delimeter overrides in your .kubecog.yaml
type Delims struct {
	Left  string `yaml:"left,omitempty"`
	Right string `yaml:"right,omitempty"`
}

// Kubecog is the type for the basic config in your .kubecog.yaml
type Kubecog struct {
	APIVersion string            `yaml:"apiVersion"`
	Kubecog    map[string]string `yaml:"kubecog"`
	Delimiters *Delims           `yaml:"delimiters,omitempty"`
}

const v1alpha1 = `v1alpha1`

type kubecogAlpha1 struct {
	APIVersion string            `yaml:"apiVersion"`
	Kubecog    map[string]string `yaml:"kubecog"`
	Delimiters *Delims           `yaml:"delimiters,omitempty"`
}

type kubecogBase struct {
	APIVersion string `yaml:"apiVersion"`
	Rest       []byte
}

func loadV1alpha1(contents []byte) (*Kubecog, error) {
	var kubecogA1 *kubecogAlpha1
	err := yaml.Unmarshal(contents, &kubecogA1)
	if err != nil {
		return nil, err
	}
	kubecog := Kubecog{
		APIVersion: kubecogA1.APIVersion,
		Kubecog:    kubecogA1.Kubecog,
		Delimiters: kubecogA1.Delimiters,
	}
	return &kubecog, nil
}

// Values is a function to get you a map of name: values.yaml
func Values() (*Kubecog, error) {
	contents, err := os.ReadFile(kubecogEnable)
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		print("No .kubecog.yaml\n")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var kubecogbase kubecogBase
	err = yaml.Unmarshal(contents, &kubecogbase)
	if err != nil {
		return nil, err
	}

	var kubecog *Kubecog
	switch kubecogbase.APIVersion {
	case v1alpha1:
		kubecog, err = loadV1alpha1(contents)
	default:
		err = ErrKubecogUnknownVersion
	}
	return kubecog, err
}
