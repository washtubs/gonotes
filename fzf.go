package main

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type FzfLister struct {
	MultiSelect bool
}

func (f *FzfLister) filterInput(input func(in io.WriteCloser)) ([]string, error) {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmdLine := "fzf"
	// TODO: why do I need this?
	cmdLine = cmdLine + " --tac"
	if f.MultiSelect {
		cmdLine = cmdLine + " -m"
	}
	cmd := exec.Command(shell, "-c", cmdLine)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, err := cmd.Output()
	if err != nil {
		return nil, errors.Errorf("Error occurred during fzf command: %s")
	}

	return strings.Split(string(result), "\n"), nil
}
