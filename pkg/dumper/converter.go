package dumper

import (
	"io"
	"os"
	"path"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
)

// DumpAndGenerateCSSfromXML generates CSS from a Chroma XML style file
// and writes it to the given writer.
func DumpAndGenerateCSSfromXML(writer io.Writer, inputFile string) {
	if !path.IsAbs(inputFile) {
		workDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		inputFile = path.Join(workDir, inputFile)
	}

	reader, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	chromaStyle := chroma.MustNewXMLStyle(reader)

	style, err := chromaStyle.Builder().Build()
	if err != nil {
		panic(err)
	}

	formatter := html.New(html.WithAllClasses(true))

	formatter.WriteCSS(writer, style)

	return
}
