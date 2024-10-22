package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Crystal() Sample {
	return Sample{
		Lexer: lexers.Get(".cr"),
		Code: `a = -1
loop do
  a += 1
  puts ((" " * (10 - a)) + ("*" * (a * 2 + 1)))
  break if a == 10
end

b = 10
loop do
  b -= 1
  puts ((" " * (10 - b)) + ("*" * (b * 2 + 1)))
  break if b == 0
end
`,
	}
}