package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Solidity() Sample {
	return Sample{
		Lexer: lexers.Get(".sol"),
		Code: `// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract HelloWorld {
    function main (string memory) public pure returns (string memory) {
        return 'Hello, World!\n';
    }
}
`,
	}
}