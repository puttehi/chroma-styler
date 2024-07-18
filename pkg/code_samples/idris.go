package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Idris() Sample {
	return Sample{
		Lexer: lexers.Get(".idr"),
		Code: `module Main

main : IO ()
main = putStrLn "Hello, World!"
`,
	}
}