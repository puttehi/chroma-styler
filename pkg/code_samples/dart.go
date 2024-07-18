package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Dart() Sample {
	return Sample{
		Lexer: lexers.Get(".dart"),
		Code: `main() {
  for (var i = 0; i < 10; i++) {
    print(" " * (10 - i) + "*" * (i * 2 + 1));
  }
  for (var i = 10; 0 <= i; i--) {
    print(" " * (10 - i) + "*" * (i * 2 + 1));
  }
}
`,
	}
}