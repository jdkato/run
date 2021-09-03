package setup

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type SetupStep interface {
	Name() string
	Run() error
}

// Fetch downloads the given resource into the provided path.
type Fetch struct {
	Link string
	Path string
}

func (f Fetch) Name() string {
	return "fetch"
}

func (f Fetch) Run() error {
	return nil
}

// A Choice represents a question posed to a user.
type Choice struct {
	Name string
	Desc string
	Link string
}

// Prompt asks a question and records the user's response.
type Prompt struct {
	Text    string
	Choices []Choice
}

func (p Prompt) Name() string {
	return "prompt"
}

func (p Prompt) Run() error {
	items := []string{}
	for _, itm := range p.Choices {
		items = append(items, itm.Name)
	}

	prompt := promptui.Select{
		Label: p.Text,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	fmt.Printf("You choose %q\n", result)
	return nil
}
