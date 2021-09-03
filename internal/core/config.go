package core

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
	"github.com/jdkato/run/internal/step"
	"github.com/mitchellh/mapstructure"
)

// A script is a series of steps for installing, configuring, and running
// arbitrary command-line tools.
//
// It's built around distinct, independently-defined step types: Setup, Update,
// Input, and Output.
type Script struct {
	Command string
	Setup   []step.Setup
}

func (s *Script) load(path string) error {
	var temp map[string]interface{}

	in, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(in, &temp)
	if err != nil {
		return err
	}

	// TODO: Error handling ...
	s.Command = temp["command"].(string)
	for _, m := range temp["setup"].([]interface{}) {
		switch m.(map[string]interface{})["type"] {
		case "fetch":
			f := step.Fetch{}
			err = mapstructure.WeakDecode(m, &f)
			s.Setup = append(s.Setup, f)
		case "prompt":
			p := step.Prompt{}
			err = mapstructure.WeakDecode(m, &p)
			s.Setup = append(s.Setup, p)
		}
	}

	return err
}

// Run the script with the provided args.
func (s *Script) Run(path string) error {
	err := s.load(path)
	if err != nil {
		return err
	}

	for _, step := range s.Setup {
		err = step.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
