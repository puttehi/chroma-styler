package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/puttehi/chroma-styler/pkg/dumper"
)

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
		dumper.DumpChromaMapping(os.Stdout)
		return
	}

	if *inputFile != "" {
		dumper.DumpAndGenerateCSSfromXML(os.Stdout, *inputFile)
		return
	}

	return
}
