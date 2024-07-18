package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Mojiscript() Sample {
	return Sample{
		Lexer: lexers.Get(".mjs"),
		Code: `import log from 'mojiscript/console/log'
import pipe from 'mojiscript/core/pipe'
import run from 'mojiscript/core/run'

const state = 'Hello, World!'

const main = pipe ([
  log
])

run ({ state, main })`,
	}
}