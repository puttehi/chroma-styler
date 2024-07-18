package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Owl_Lisp() Sample {
	return Sample{
		Lexer: lexers.Get(".scm"),
		Code: `(λ (args) (print "Hello, World!"))
`,
	}
}