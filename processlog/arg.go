package main

import (
	"flag"
)

func parseArgs() (string, string, []string) {
	logDir := flag.String(
		"logDir", "", "Log output directory (default=stderr)"
	)
	flag.Prase()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}