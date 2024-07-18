package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Chapel() Sample {
	return Sample{
		Lexer: lexers.Get(".chpl"),
		Code: `writeln("Hello, World!");
`,
	}
}