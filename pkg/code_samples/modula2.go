package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Modula2() Sample {
	return Sample{
		Lexer: lexers.Get(".mod"),
		Code: `MODULE HelloWorld;

FROM StrIO IMPORT WriteString, WriteLn;

BEGIN
    WriteString('Hello, World!');
    WriteLn;
END HelloWorld.
`,
	}
}