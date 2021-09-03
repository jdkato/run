package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jdkato/run/internal/cli"
	"github.com/jdkato/run/internal/core"
	"github.com/jdkato/run/internal/setup"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// version is set during the release build process.
var version = "master"

func doRun(args []string) (core.Script, error) {
	var script core.Script
	var temp map[string]interface{}

	in, err := ioutil.ReadFile(args[0])
	if err != nil {
		return script, err
	}

	err = yaml.Unmarshal(in, &temp)

	// TODO: Error handling ...
	script.Command = temp["command"].(string)

	for _, m := range temp["setup"].([]interface{}) {
		switch m.(map[string]interface{})["type"] {
		case "fetch":
			f := setup.Fetch{}
			err = mapstructure.WeakDecode(m, &f)
			script.Setup = append(script.Setup, f)
		case "prompt":
			p := setup.Prompt{}
			err = mapstructure.WeakDecode(m, &p)
			script.Setup = append(script.Setup, p)
		}
	}

	return script, err
}

func main() {
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

	script, err := doRun(args)
	if err != nil {
		panic(err)
	}

	for _, step := range script.Setup {
		if step.Name() == "prompt" {
			v := step.(setup.Prompt)

			items := []string{}
			for _, itm := range v.Choices {
				items = append(items, itm.Name)
			}

			prompt := promptui.Select{
				Label: v.Text,
				Items: items,
			}

			_, result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %q\n", result)
		}
	}

	os.Exit(0)
}
