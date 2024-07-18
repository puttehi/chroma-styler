package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Io() Sample {
	return Sample{
		Lexer: lexers.Get(".io"),
		Code: `"Hello, World!" println
`,
	}
}