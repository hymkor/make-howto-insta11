package main

import (
	"fmt"
	"os"

	"github.com/hymkor/make-howto-insta11/internal/gitdir"
)

const theTemplate = `[![GoDev](https://pkg.go.dev/badge/github.com/%[2]s/%[3]s)](https://pkg.go.dev/github.com/%[2]s/%[3]s)

-----

Install
-------

Download the binary package from [Releases](https://github.com/%[2]s/%[3]s/releases) and extract the executable.

### for scoop-installer

%[1]s
scoop install https://raw.githubusercontent.com/%[2]s/%[3]s/master/%[4]s.json
%[1]s

or

%[1]s
scoop bucket add %[2]s https://github.com/%[2]s/scoop-bucket
scoop install %[4]s
%[1]s
`

func mains() error {
	user, repo, err := gitdir.GetNameAndRepo()
	if err != nil {
		return err
	}
	fmt.Printf(theTemplate, "```", user, repo, repo)
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
