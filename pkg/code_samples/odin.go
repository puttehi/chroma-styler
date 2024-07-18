package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Odin() Sample {
	return Sample{
		Lexer: lexers.Get(".odin"),
		Code: `package main

import "core:fmt"
import "core:strconv"
import "core:os"

main :: proc() {
    if len(os.args) != 2 {
        usage()
        return
    }
    n, ok := strconv.parse_int(os.args[1])
    if !ok {
        usage()
        return
    }
    evenodd(n)
}

evenodd :: proc(n: int) {
    if n % 2 == 0 {
        fmt.println("Even")
    } else {
        fmt.println("Odd")
    }
}

usage :: proc() {
    fmt.println("Usage: please input a number")
}`,
	}
}