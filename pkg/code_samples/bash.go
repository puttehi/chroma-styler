package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Bash() Sample {
	return Sample{
		Lexer: lexers.Get(".sh"),
		Code: `#!/bin/bash

for i in {0..9}; do
	printf " %.0s" $(seq 1 $(( 10 - $i )))
	printf "*%.0s" $(seq 1 $(( $i * 2 + 1 )))
	echo
done

printf "*%.0s" {1..21}
echo

for i in {9..0}; do
	printf " %.0s" $(seq 1 $(( 10 - $i )))
	printf "*%.0s" $(seq 1 $(( $i * 2 + 1 )))
	echo
done
`,
	}
}