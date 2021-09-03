package core

import (
	"bytes"

	"github.com/alecthomas/chroma/quick"
)

// Highlight performs syntax-highlighting on the given string.
func Highlight(s string) string {
	buf := new(bytes.Buffer)
	_ = quick.Highlight(buf, s, "yaml", "terminal256", "manni")
	return buf.String()
}
