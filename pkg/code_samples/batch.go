package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Batch() Sample {
	return Sample{
		Lexer: lexers.Get(".cmd"),
		Code: `@echo off
echo Hello, World!
`,
	}
}