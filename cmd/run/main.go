package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jdkato/run/internal/cli"
	"github.com/jdkato/run/internal/core"
)

// version is set during the release build process.
var version = "master"

func main() {
	var script core.Script

	v := flag.Bool("v", false, "prints the current version")
	flag.Parse()

	if *v {
		fmt.Println(version)
		os.Exit(0)
	}

	args := flag.Args()
	argc := len(args)

	if argc == 0 {
		cli.PrintIntro()
	} else {
		cmd, exists := cli.Actions[args[0]]
		if exists {
			if err := cmd(args[1:]); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}

	err := script.Run(args[0])
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
