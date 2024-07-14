package viewer

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/puttehi/chroma-styler/pkg/code_samples"
)

type Server struct {
	router *gin.Engine
	port   int
}

type TemplateData struct {
	CodeBlocks []string
}

func New(port int, cssFilePath string) Server {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery()) // anti-panic

	if !path.IsAbs(cssFilePath) {
		workDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		cssFilePath = path.Join(workDir, cssFilePath)
	}

	r.StaticFile("syntax.css", cssFilePath) // User CSS

	fmt.Printf(">>>>>>>>>> \n/syntax.css => %s \n>>>>>>>>>>\n", cssFilePath)

	renderer := NewRenderer()

	tpl, err := template.New("index").Parse(Template)

	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		rendered, err := renderer.Format(code_samples.Go())
		if err != nil {
			panic(err)
		}
		err = tpl.Execute(c.Writer, TemplateData{
			CodeBlocks: []string{rendered, rendered},
		})
		if err != nil {
			panic(err)
		}
	})

	return Server{
		router: r,
		port:   port,
	}
}

func (s Server) Serve() {
	fmt.Printf("**********\nStarting on http://localhost:%d\n**********\n", s.port)

	s.router.Run(fmt.Sprintf(":%d", s.port))
}
