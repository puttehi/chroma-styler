package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Terra() Sample {
	return Sample{
		Lexer: lexers.Get(".t"),
		Code: `print("Hello, World!")
`,
	}
}