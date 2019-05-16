package main

import (
	"io"
	"os"
	"os/exec"
)

func excution(commandName string, args []string, stdout, stderr io.Writer)(*os.processState, error) {
	cmd := exec.Command(commandName, args...)
	childStdout, _ := cmd.StdoutPips()
	childStderr, _ := cmd.StderrPips()

	go io.Copy(stdout, childStdout)
	go io.Copy(stderr, childStderr)

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return cmd.ProcessState, nil
}