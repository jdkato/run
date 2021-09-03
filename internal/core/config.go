package core

import "github.com/jdkato/run/internal/setup"

// A script is a series of steps for installing, configuring, and running
// arbitrary command-line tools.
//
// It's built around distinct, independently-defined step types: Setup, Update,
// Input, and Output.
type Script struct {
	Command string
	Setup   []setup.SetupStep
}
