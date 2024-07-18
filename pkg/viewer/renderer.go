package viewer

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/puttehi/chroma-styler/pkg/code_samples"
)

type ChromaRenderer struct {
	formatter *html.Formatter
}

func NewRenderer() ChromaRenderer {
	return ChromaRenderer{
		formatter: html.New(html.WithClasses(true)),
	}
}

func (cr ChromaRenderer) Format(sample code_samples.Sample) (string, error) {
	fmt.Printf("Formatting %s\n", sample.Lexer.Config().Name)
	it, err := sample.Lexer.Tokenise(nil, sample.Code)
	if err != nil {
		return sample.Code, err
	}

	var buf bytes.Buffer
	err = cr.formatter.Format(&buf, styles.Fallback, it)
	if err != nil {
		return sample.Code, err
	}

	return buf.String(), nil
}
