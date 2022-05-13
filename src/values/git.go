package values

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"regexp"
)

// gitURLEnv is the name of the environment variable controlling where
// the git repo is for cog
const gitURLEnv = `COG_GIT_URL`

// gitURLref is the name of the environment variable controlling which
// git repo we are using for cog
const gitRefEnv = `COG_GIT_REF`

func gitURL() string {
	if url, pathpresent := os.LookupEnv(gitURLEnv); pathpresent {
		return url
	}
	return ``
}

func gitRef() plumbing.ReferenceName {
	if ref, pathpresent := os.LookupEnv(gitRefEnv); pathpresent {
		return plumbing.ReferenceName(ref)
	}
	return ``
}

var sanitiseRe = regexp.MustCompile(`[:/\\<>\|]`)

func sanitiseToPath(in string) string {
	return sanitiseRe.ReplaceAllString(in, `_`)
}

func gitDirectory() string {
	return os.TempDir() + `/` + sanitiseToPath(gitURL()) + sanitiseToPath(string(gitRef()))
}

func fromGit() (string, error) {
	url := gitURL()
	if url == `` {
		return ``, nil
	}
	ref := gitRef()
	directory := gitDirectory()
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		ReferenceName:     ref,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	if err != nil && err != git.ErrRepositoryAlreadyExists {
		return ``, err
	}
	if err == git.ErrRepositoryAlreadyExists {
		r, err = git.PlainOpen(directory)
		if err != nil {
			return ``, err
		}
	}
	w, err := r.Worktree()
	if err != nil {
		return ``, err
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return ``, err
	}

	return directory, nil
}
