package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Verilog() Sample {
	return Sample{
		Lexer: lexers.Get(".v"),
		Code: `module main;
  initial
    begin
      $display("Hello, World!");
      $finish(0);
    end
endmodule
`,
	}
}