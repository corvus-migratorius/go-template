package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

const VERSION = "0.0.1"

type Args struct {
	Version bool
}

func parseArgs() Args {
	parser := argparse.NewParser(fmt.Sprintf("app (v. %s)", VERSION), "Template app")
	version := parser.Flag("V", "version", &argparse.Options{Required: false, Help: "Display program version and exit", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	return Args{
		Version: *version,
	}
}

func main() {
	_ = parseArgs()
}
