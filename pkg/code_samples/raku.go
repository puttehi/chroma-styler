package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Raku() Sample {
	return Sample{
		Lexer: lexers.Get(".raku"),
		Code: `say 'Hello, World!'
`,
	}
}