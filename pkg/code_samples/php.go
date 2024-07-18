package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Php() Sample {
	return Sample{
		Lexer: lexers.Get(".php"),
		Code: `<?php

for ($i = 0; $i < 10; $i++) {
    echo str_repeat(' ', 10 - $i) . str_repeat('*', $i * 2 + 1) . "\n";
}

for ($i = 10;-1 < $i; $i--) {
    echo str_repeat(' ', 10 - $i) . str_repeat('*', $i * 2 + 1) . "\n";
}
`,
	}
}