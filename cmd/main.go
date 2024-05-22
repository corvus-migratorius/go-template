package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var Version = "development"

// Args represents parsed CLI arguments
type args struct {
	Version bool
}

func parseArgs() args {
	parser := argparse.NewParser(fmt.Sprintf("app %s", Version), "Template app")
	showVersion := parser.Flag("V", "version", &argparse.Options{Required: false, Help: "Display program version and exit", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	return args{}
}

func main() {
	_ = parseArgs()
}
