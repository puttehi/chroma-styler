package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func V() Sample {
	return Sample{
		Lexer: lexers.Get(".v"),
		Code: `module main

fn main(){
    println('Hello, World!')
}
`,
	}
}