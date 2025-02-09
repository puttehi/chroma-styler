package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Fortran() Sample {
	return Sample{
		Lexer: lexers.Get(".f95"),
		Code: `program Baklava
    do i = 0, 10, 1
        print '(10a)', repeat (" ", (10 - i)), repeat ("*", (i * 2 + 1))
    end do
    do i = 9, 0, -1
        print '(10a)', repeat (" ", (10 - i)), repeat ("*", (i * 2 + 1))
    end do
end program Baklava
`,
	}
}