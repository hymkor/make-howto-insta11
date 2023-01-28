package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
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

func quote(args []string, f func(string) error) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	r, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer r.Close()
	cmd.Start()

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		//println(sc.Text())
		if err := f(sc.Text()); err != nil {
			return err
		}
	}
	return nil
}

func listUpRemoteBranch() ([]string, error) {
	branches := []string{}
	quote([]string{"git", "remote", "show"}, func(line string) error {
		branches = append(branches, strings.TrimSpace(line))
		return nil
	})
	return branches, nil
}

var rxURL = regexp.MustCompile(`Push +URL: \w+@github.com:([\w-]+)/([\w-]+).git`)

func getNameAndRepo() (string, string, error) {
	branch, err := listUpRemoteBranch()
	if err != nil {
		return "", "", err
	}
	if len(branch) < 1 {
		return "", "", errors.New("remote branch not found")
	}
	var user, repo string
	quote([]string{"git", "remote", "show", "-n", branch[0]}, func(line string) error {
		m := rxURL.FindStringSubmatch(line)
		if m != nil {
			user = m[1]
			repo = m[2]
			return io.EOF
		}
		return nil
	})
	return user, repo, nil
}

func mains() error {
	user, repo, err := getNameAndRepo()
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
