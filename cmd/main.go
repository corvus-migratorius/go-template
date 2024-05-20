package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

// VERSION indicates program version
const version = "0.0.1"

// Args represents parsed CLI arguments
type args struct {
	Version bool
}

func parseArgs() args {
	parser := argparse.NewParser(fmt.Sprintf("app (v. %s)", version), "Template app")
	v := parser.Flag("V", "version", &argparse.Options{Required: false, Help: "Display program version and exit", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *v {
		fmt.Println(version)
		os.Exit(0)
	}

	return args{
		Version: *v,
	}
}

func main() {
	_ = parseArgs()
}
