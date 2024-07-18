package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Hack() Sample {
	return Sample{
		Lexer: lexers.Get(".hh"),
		Code: `<<__EntryPoint>>
function main(): void {
    echo "Hello, World!";
}
`,
	}
}