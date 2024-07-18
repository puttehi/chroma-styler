package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Coffeescript() Sample {
	return Sample{
		Lexer: lexers.Get(".coffee"),
		Code: `for i in [0...10]
  pattern = " ".repeat (10 - i)
  pattern += "*".repeat (i * 2 + 1)
  console.log pattern

for i in [10..0]
  pattern = " ".repeat (10 - i)
  pattern += "*".repeat (i * 2 + 1)
  console.log pattern`,
	}
}