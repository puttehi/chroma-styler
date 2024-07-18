package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Eta() Sample {
	return Sample{
		Lexer: lexers.Get(".hs"),
		Code: `module Main where
main :: IO ()
main = putStrLn "Hello, World!"
`,
	}
}