package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Ferret() Sample {
	return Sample{
		Lexer: lexers.Get(".clj"),
		Code: `(do
  (println "Hello, World!"))
`,
	}
}