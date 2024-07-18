package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func C() Sample {
	return Sample{
		Lexer: lexers.Get(".c"),
		Code: `#include "stdio.h"

int main (void)
{

  for (int i = 0; i < 10; i++)
  {
    printf ("%.*s", (10 - i), "                                 ");
    printf ("%.*s", (i * 2 + 1), "******************************");
    printf ("\n");
  }

  for (int i = 10; -1 < i; i--)
  {
    printf ("%.*s", (10 - i), "                                 ");
    printf ("%.*s", (i * 2 + 1), "******************************");
    printf ("\n");
  }

  return 0;

}
`,
	}
}