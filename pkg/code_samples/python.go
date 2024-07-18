package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Python() Sample {
	return Sample{
		Lexer: lexers.Get(".py"),
		Code: `for i in range(0, 10, 1):
    print((" " * (10 - i)) + ("*" * (i * 2 + 1)))

for i in range(10, -1, -1): 
    print((" " * (10 - i)) + ("*" * (i * 2 + 1)))
`,
	}
}