package dumper

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/markkurossi/tabulate"
)

// DumpChromaMapping generates a table of
//
// * Chroma style keywords
//
// * Their CSS class mapping (1:1)
//
// * Example regex used by the Chroma lexer for those highlights
//
// And writes them to the given writer.
func DumpChromaMapping(writer io.Writer) {
	// Sort the keys by the underlying TokenType int iota
	keys := make([]int, 0, len(chroma.StandardTypes))
	for k := range chroma.StandardTypes {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	// Get example regexes from Go lexer
	reLexer, ok := lexers.Go.(*chroma.RegexLexer)
	if !ok {
		panic("Go lexer is not a RegexLexer")
	}
	lexerRules, err := reLexer.Rules()
	if err != nil {
		panic(err)
	}
	rulesArr, ok := lexerRules["root"]
	if !ok {
		panic("Missing 'root' key in lexer rules")
	}
	rulesMap := make(map[string]string) // Type => Pattern
	for _, r := range rulesArr {
		// TODO: One Type might have multiple regexes, should tabulate with newlines to see each option.
		typeStr := fmt.Sprintf("%v", r.Type) // It's not stringable but this works through reflection it seems
		rulesMap[typeStr] = r.Pattern
	}

	// Format to raw rows
	var lines []string
	const splitToken = "**FOOBAR**" // Regexes might have funky stuff which break later row building
	for _, k := range keys {
		tt := chroma.TokenType(k)
		rePattern := rulesMap[tt.String()]
		lines = append(lines, fmt.Sprintf("%s%s%s%s%s", tt.String(), splitToken, chroma.StandardTypes[tt], splitToken, rePattern))
	}

	// Tabulate for a pretty output
	tab := tabulate.New(tabulate.Simple)
	tab.Header("Chroma style")
	tab.Header("CSS class name")
	tab.Header("Chroma Regex (for Go)")
	for _, line := range lines {
		row := tab.Row()
		for _, col := range strings.Split(line, splitToken) {
			row.Column(col)
		}
	}

	tab.Print(writer)

	return
}
