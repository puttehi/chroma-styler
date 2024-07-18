package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Zig() Sample {
	return Sample{
		Lexer: lexers.Get(".zig"),
		Code: `const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.writeAll("Hello, World!\n");
}
`,
	}
}