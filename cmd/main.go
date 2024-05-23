package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var version = "development"

type appArgs struct{}

func parseArgs() appArgs {
	parser := argparse.NewParser(fmt.Sprintf("template %s", version), "Template template")
	showVersion := parser.Flag(
		"V", "version",
		&argparse.Options{Required: false, Help: "Display program version and exit", Default: false},
	)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	return appArgs{}
}

func main() {
	_ = parseArgs()
}
