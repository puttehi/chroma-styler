package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Ballerina() Sample {
	return Sample{
		Lexer: lexers.Get(".bal"),
		Code: `import ballerina/io;
public function main(string... args) {
    io:println("Hello, World!");
}
`,
	}
}