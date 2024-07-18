package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Scilab() Sample {
	return Sample{
		Lexer: lexers.Get(".sce"),
		Code: `mprintf('%s', 'Hello, World!')
`,
	}
}