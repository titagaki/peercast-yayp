package main

import (
	"html/template"
	"io"
	"log"
	"peercast-yayp/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	//e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.tmpl")),
	}
	e.Renderer = renderer

	e.Static("/static", "assets")
	e.File("/favicon.ico", "assets/favicon.ico")
	e.File("/robots.txt", "assets/robots.txt")

	e.GET("/", handler.TopPage())
	e.GET("/getgmt", handler.GetGmt())
	e.GET("/index.txt", handler.IndexTxt())

	e.Logger.Fatal(e.Start(":8000"))
}
