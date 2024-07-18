package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Scheme() Sample {
	return Sample{
		Lexer: lexers.Get(".scm"),
		Code: `(display "Hello, World!")
`,
	}
}