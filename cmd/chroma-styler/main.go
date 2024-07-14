package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/puttehi/chroma-styler/pkg/dumper"
	"github.com/puttehi/chroma-styler/pkg/viewer"
)

func main() {
	const (
		serveViewDefault = false
		dumpOnlyDefault  = false
		inputFileDefault = ""
	)

	serveView := flag.Bool("viewer", serveViewDefault, "If given, starts a HTTP server for viewing several code samples with Chroma highlighting.")
	dumpOnly := flag.Bool("dump-types", dumpOnlyDefault, "If given, only dumps the Chroma Token -> CSS class mapping for building your own styles.")
	inputFile := flag.String("input", inputFileDefault, "Input Chroma XML file to generate CSS from. Output will be printed to stdout.")
	flag.Parse()

	if *serveView == serveViewDefault && *dumpOnly == dumpOnlyDefault && *inputFile == inputFileDefault {
		fmt.Printf("No flags given.\n")
		flag.Usage()
		os.Exit(2)
	}

	if *serveView {
		server := viewer.New(3000)
		server.Serve()
		os.Exit(0)
	}

	if *dumpOnly {
		dumper.DumpChromaMapping(os.Stdout)
		os.Exit(0)
	}

	if *inputFile != "" {
		dumper.DumpAndGenerateCSSfromXML(os.Stdout, *inputFile)
		os.Exit(0)
	}

	return
}
