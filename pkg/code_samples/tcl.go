package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Tcl() Sample {
	return Sample{
		Lexer: lexers.Get(".tcl"),
		Code: `puts "Hello, World!"
`,
	}
}