package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Chaiscript() Sample {
	return Sample{
		Lexer: lexers.Get(".chai"),
		Code: `print("Hello, World!")
`,
	}
}