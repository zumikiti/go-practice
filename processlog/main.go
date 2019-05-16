package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logDir, commandName, args := parseArgs()
	stdout, stderr, err := initOut(logDir, commandName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create file: %v\n", err)
		os.Exit(0)
	}

	startTime := time.Now()
	ps, err := execution.(commandName, args, stdout, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "Cant exec command %s", commandName)
		os.Exit(0)
	}
	exitTime := time.Now();
	fmt.Fprintf(stdout, ps.String())
	fmt.Fprintf(stdout, "wall clock time=%v user time=%v\n", exitTime.Sub(startTime), ps.SystemTime(), ps.SystemTime(), ps.UserTime())
}