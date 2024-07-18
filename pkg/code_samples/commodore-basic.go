package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Commodore_Basic() Sample {
	return Sample{
		Lexer: lexers.Get(".bas"),
		Code: `10 FOR I = -10 TO 10
20     NSP% = ABS(I)
30     SP$ = ""
40     IF NSP% = 0 GOTO 80
50     FOR J = 1 TO NSP%
60         SP$ = SP$ + " "
70     NEXT J
80     ST$ = ""
90     FOR J = 1 TO 21 - 2 * NSP%
100         ST$ = ST$ + "*"
110     NEXT J
120     PRINT SP$ + ST$
130 NEXT I
`,
	}
}