package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Groovy() Sample {
	return Sample{
		Lexer: lexers.Get(".groovy"),
		Code: `(0..10).each{ index ->
    println "${' ' * (10 - index)}${'*' * (index * 2 + 1)}"
}
(9..0).each{ index ->
    println "${' ' * (10 - index)}${'*' * (index * 2 + 1)}"
}`,
	}
}