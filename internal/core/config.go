package core

import (
	"gopkg.in/yaml.v3"
)

// A script is a series of steps for installing, configuring, and running
// arbitrary command-line tools.
//
// It's built around distinct, independently-defined step types: Setup, Update,
// Input, and Output.
type Script struct {
	Setup []Setup
}

func init() {
	var v struct {
		A int
		B string
	}
	v.A = 1
	v.B = "hello"
	_, _ = yaml.Marshal(v)
}
