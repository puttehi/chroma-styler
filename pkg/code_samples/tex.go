package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Tex() Sample {
	return Sample{
		Lexer: lexers.Get(".tex"),
		Code: `%&pdftex

\countdef\counter=1
\countdef\threecounter=2
\countdef\fivecounter=3
\counter=1
\threecounter=1
\fivecounter=1

\loop
    \def\printnum{1}
    \noindent
    \ifnum \threecounter>2
        Fizz%
        \threecounter=0
        \def\printnum{0}
    \fi
    \ifnum \fivecounter>4
        Buzz%
        \fivecounter=0
        \def\printnum{0}
    \fi
    \ifnum\printnum=1
        % print the counter variable
        \the\counter%
    \fi
    \hfil \break
    \advance \counter 1
    \advance \threecounter 1
    \advance \fivecounter 1
\ifnum \counter<101
\repeat

\end
`,
	}
}