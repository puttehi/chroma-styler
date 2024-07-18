package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Rexx() Sample {
	return Sample{
		Lexer: lexers.Get(".rexx"),
		Code: `/* ARG with source string named in REXX program invocation */
arg number
If (DATATYPE(number, 'W') == 0) then signal err
if (number // 2 == 0) then
	say "Even"
else
	say "Odd"
;exit

Err:
say 'Usage: please input a number'; exit
`,
	}
}