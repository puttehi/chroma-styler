package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Ruby() Sample {
	return Sample{
		Lexer: lexers.Get(".rb"),
		Code: `10.times {|n|
    print " " * (10 - n)
    puts "*" * (2 * n + 1)
}

for i in 10.downto(0) do
  puts ((" " * (10 - i)) + ("*" * (i * 2 + 1)))
end
`,
	}
}