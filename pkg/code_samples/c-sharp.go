package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func C_Sharp() Sample {
	return Sample{
		Lexer: lexers.Get(".cs"),
		Code: `using System;

class CSharp
{

    static void Main (string[] args)
    {

        for (SByte i = 0; i < 10; i++)
            Console.WriteLine (
                new string (' ', (10 - i)) + new string ('*', (i * 2 + 1))
            );

        for (SByte i = 10; -1 < i; i--)
            Console.WriteLine (
                new string (' ', (10 - i)) + new string ('*', (i * 2 + 1))
            );

    }

}
`,
	}
}