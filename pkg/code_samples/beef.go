package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Beef() Sample {
	return Sample{
		Lexer: lexers.Get(".bf"),
		Code: `using System;

namespace Baklava;

class Program
{
    public static int Main(String[] args)
    {
        for (int i in -10...10)
        {
            int num_spaces = Math.Abs(i);
            int num_stars = 21 - 2 * num_spaces;
            String line = scope String();
            line.Append(' ', num_spaces);
            line.Append('*', num_stars);
            Console.WriteLine(line);
        }
        return 0;
    }
}
`,
	}
}