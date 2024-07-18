package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Javascript() Sample {
	return Sample{
		Lexer: lexers.Get(".js"),
		Code: `for (var i = 0; i < 10; i++)
    console.log (
        " ".repeat (10 - i) + "*".repeat (i * 2 + 1)
    );

for (var i = 10; -1 < i; i--)
    console.log (
        " ".repeat (10 - i) + "*".repeat (i * 2 + 1)
    );
`,
	}
}