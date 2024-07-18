//go:generate go run ../../cmd/snippet-generator/main.go ../../data/sample-programs/archive
package code_samples

import "github.com/alecthomas/chroma/v2"

type Sample struct {
	Code  string
	Lexer chroma.Lexer
}
