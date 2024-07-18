package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Pony() Sample {
	return Sample{
		Lexer: lexers.Get(".pony"),
		Code: `actor Main
  new create(env: Env) =>
    env.out.print("Hello, World!")
`,
	}
}