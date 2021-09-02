package core

// Fetch downloads the given resource into the provided path.
type Fetch struct {
	Link string
	Path string
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

// Setup represents a configuration script for a given command-line tool.
//
// It handles tasks such as layout scaffolding, file creation, and resource
// downloading.
type Setup struct {
	Fetches []Fetch
	Prompts []Prompt
}
