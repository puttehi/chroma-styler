package viewer

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port   int
}

type TemplateData struct {
	CodeBlocks []string
}

func New(port int) Server {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery()) // anti-panic

	tpl, err := template.New("index").Parse(Template)
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		err := tpl.Execute(c.Writer, TemplateData{
			CodeBlocks: []string{"first", "second"},
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
	fmt.Printf("\nStarting on http://localhost:%d\n*****\n\n", s.port)

	s.router.Run(fmt.Sprintf(":%d", s.port))
}
