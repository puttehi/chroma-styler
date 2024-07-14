package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Go() Sample {
	return Sample{
		Lexer: lexers.Go,
		Code: `package main

import (
	"errors"
	"fmt"
)

type Foo struct {
	concrete string
	function func(param int) error
}

type Bar interface {
	Baz() (Foo, error)
}

func dependencyInjection(bar Bar) (Foo, error) {
	return bar.Baz()
}

func (f Foo) Baz() (Foo, error) {
	fmt.Printf("%s %d", "foo", 123)
	return Foo{}, errors.New("bar")
}

func main() {
	_, _ = Foo{}.Baz()
	return
}`,
	}
}
