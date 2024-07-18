package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Swift() Sample {
	return Sample{
		Lexer: lexers.Get(".swift"),
		Code: `func baklava(maxWidth: Int) -> Void {
    for number in 0...maxWidth {
        print(String(repeatElement(" ", count:maxWidth-number)) + String(repeatElement("*", count:number*2+1)))
    }
    for number in (0...maxWidth-1).reversed() {
        print(String(repeatElement(" ", count:maxWidth-number)) + String(repeatElement("*", count:number*2+1)))
    }
}
baklava(maxWidth: 10)
`,
	}
}