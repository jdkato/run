package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jdkato/run/internal/cli"
)

// version is set during the release build process.
var version = "master"

func main() {
	v := flag.Bool("v", false, "prints current version")
	flag.Parse()

	if *v {
		fmt.Println("vale version " + version)
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

	os.Exit(0)
}
