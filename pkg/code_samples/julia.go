package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Julia() Sample {
	return Sample{
		Lexer: lexers.Get(".jl"),
		Code: `#!/usr/bin/julia

function main()
    for i = 0:10
        print(" "^(10 - i))
        println("*"^(i * 2 + 1))
    end 

    for i = 9:-1:0
        print(" "^(10 - i))
        println("*"^(i * 2 + 1))
    end
end

main()
`,
	}
}