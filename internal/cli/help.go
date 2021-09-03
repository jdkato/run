package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/jdkato/run/internal/core"
	"github.com/logrusorgru/aurora/v3"
	"github.com/olekukonko/tablewriter"
)

var exampleConfig = `# $ run vale.yml
	---
	language: python
	sudo: false
	python:
	  - 2.7
	  - 3.4
	script: py.test -v`

var intro = fmt.Sprintf(`%s

%s:	%s

%s is a tool for running other command-line tools: create interactive prompts,
collect specific inputs, manage configuration files, and much more!

To get started, you'll need a script.yml file.

%s:

	%s

See %s for more setup information.`,
	aurora.Bold("run - A meta command-line tool."),
	aurora.Bold("Usage"),
	aurora.Faint("run [options] script.yml"),
	aurora.Italic("run"),
	aurora.Bold("Example"),
	core.Highlight(exampleConfig),
	aurora.Underline("https://docs.errata.ai/vale/about"))

var info = fmt.Sprintf(`%s

(Or use %s for a listing of all CLI options.)`,
	intro,
	aurora.Faint("run --help"))

// PrintIntro shows basic usage / gettting started info.
func PrintIntro() {
	fmt.Println(info)
	os.Exit(0)
}

func init() {
	flag.Usage = func() {
		fmt.Println(intro)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetAutoWrapText(false)

		fmt.Println(aurora.Bold("\nFlags:"))
		flag.VisitAll(func(f *flag.Flag) {
			table.Append([]string{"--" + f.Name, f.Usage})
		})

		table.Render()
		table.ClearRows()

		fmt.Println(aurora.Bold("Commands:"))
		for cmd, use := range commandInfo {
			table.Append([]string{cmd, use})
		}
		table.Render()

		os.Exit(0)
	}
}
