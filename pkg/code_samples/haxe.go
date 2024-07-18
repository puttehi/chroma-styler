package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Haxe() Sample {
	return Sample{
		Lexer: lexers.Get(".hx"),
		Code: `class HelloWorld {
    static public function main() {
        trace("Hello, World!");
    }
}
`,
	}
}