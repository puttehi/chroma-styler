package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func Perl() Sample {
	return Sample{
		Lexer: lexers.Get(".pl"),
		Code: `#!/usr/bin/env perl
use strict;
use warnings;

my $size = 10;

for my $i (1..$size){
    print " "x($size + 1 - $i), "*"x($i*2 - 1), "\n";
}

for my $j (0..$size){
    print " "x($j), "*"x($size*2 - $j*2 + 1), "\n";
}`,
	}
}