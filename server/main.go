package main

import (
	"html/template"
	"io"
	"peercast-yayp/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.tmpl")),
	}
	e.Renderer = renderer

	e.GET("/", handler.TopPage())
	e.GET("/index.txt", handler.IndexTxt())
	e.GET("/getgmt.*", handler.ChannelStatistics())
	e.GET("/chat.*", handler.Chat())

	e.Static("/*", "public")

	e.Logger.Fatal(e.Start(":8000"))
}
