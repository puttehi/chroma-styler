package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Ocaml() Sample {
	return Sample{
		Lexer: lexers.Get(".ml"),
		Code: `print_string "Hello, World!\n";;
`,
	}
}