package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	// "strconv"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	tab "github.com/markkurossi/tabulate"
)

func dumpChromaMapping() {
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
	tab := tab.New(tab.Simple)
	tab.Header("Chroma style")
	tab.Header("CSS class name")
	tab.Header("Chroma Regex (for Go)")
	for _, line := range lines {
		row := tab.Row()
		for _, col := range strings.Split(line, splitToken) {
			row.Column(col)
		}
	}

	tab.Print(os.Stdout)

	return
}

func dumpAndGenerateCSSfromXML(inputFile string) {
	// Root directory is current working directory
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	absPath := path.Join(workDir, inputFile)

	reader, err := os.Open(absPath)

	if err != nil {
		panic(err)
	}

	chromaStyle := chroma.MustNewXMLStyle(reader)

	style, err := chromaStyle.Builder().Build()
	if err != nil {
		panic(err)
	}

	formatter := html.New(html.WithAllClasses(true))

	formatter.WriteCSS(os.Stdout, style)

	return
}

func main() {
	const (
		dumpOnlyDefault  = false
		inputFileDefault = ""
	)

	dumpOnly := flag.Bool("dump-types", dumpOnlyDefault, "If given, only dumps the Chroma Token -> CSS class mapping for building your own styles.")
	inputFile := flag.String("input", inputFileDefault, "Input Chroma XML file to generate CSS from. Output will be printed to stdout.")
	flag.Parse()

	if *dumpOnly == dumpOnlyDefault && *inputFile == inputFileDefault {
		fmt.Printf("No flags given.\n")
		flag.Usage()
		return
	}

	if *dumpOnly {
		dumpChromaMapping()
		return
	}

	if *inputFile != "" {
		dumpAndGenerateCSSfromXML(*inputFile)
		return
	}

	return
}
