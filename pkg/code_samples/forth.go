package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Forth() Sample {
	return Sample{
		Lexer: lexers.Get(".fth"),
		Code: `.( Hello, World!) CR
bye
`,
	}
}