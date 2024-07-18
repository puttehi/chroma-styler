package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Smalltalk() Sample {
	return Sample{
		Lexer: lexers.Get(".st"),
		Code: `Transcript show: 'Hello, World!'.
`,
	}
}