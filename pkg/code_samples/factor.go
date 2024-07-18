package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Factor() Sample {
	return Sample{
		Lexer: lexers.Get(".factor"),
		Code: `USE: io

"Hello, World!" print
`,
	}
}