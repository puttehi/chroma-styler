package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Gerbil() Sample {
	return Sample{
		Lexer: lexers.Get(".ss"),
		Code: `(export main)
(def (main . args)
  (displayln "Hello, World!")
)
`,
	}
}