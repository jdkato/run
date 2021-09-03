package setup

type SetupStep interface {
	Name() string
}

// Fetch downloads the given resource into the provided path.
type Fetch struct {
	Link string
	Path string
}

func (f Fetch) Name() string {
	return "fetch"
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
