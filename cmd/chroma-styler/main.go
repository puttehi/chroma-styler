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
		serveViewDefault   = false
		cssFilePathDefault = ""
		dumpOnlyDefault    = false
		inputFileDefault   = ""
	)

	serveView := flag.Bool("viewer", serveViewDefault, "If given, starts a HTTP server for viewing several code samples with Chroma highlighting. Pass in syntax.css with -css.")
	cssFilePath := flag.String("css", cssFilePathDefault, "Which CSS file to use for syntax highlighting? Must be Chroma class-compliant (use -input first to create one).")
	dumpOnly := flag.Bool("dump-types", dumpOnlyDefault, "If given, only dumps the Chroma Token -> CSS class mapping for building your own styles.")
	inputFile := flag.String("input", inputFileDefault, "Input Chroma XML file to generate CSS from. Output will be printed to stdout.")
	flag.Parse()

	if *serveView == serveViewDefault && *cssFilePath == cssFilePathDefault && *dumpOnly == dumpOnlyDefault && *inputFile == inputFileDefault {
		fmt.Printf("No flags given.\n")
		flag.Usage()
		os.Exit(2)
	}

	if *serveView {
		if *cssFilePath == "" {
			fmt.Printf("Did not give a value for the argument -css. Cannot load styles.")
			os.Exit(2)
		}
		server := viewer.New(3000, *cssFilePath)
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
