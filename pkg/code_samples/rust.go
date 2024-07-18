package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Rust() Sample {
	return Sample{
		Lexer: lexers.Get(".rs"),
		Code: `fn main() {
    for i in -10i32..11i32 {
        let num_spaces: usize = i.abs() as usize;
        println!("{}{}", " ".repeat(num_spaces), "*".repeat(21 - 2 * num_spaces));
    }
}
`,
	}
}