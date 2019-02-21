package main

import (
	"html/template"
	"io"
	"peercast-yayp/handler"
	"peercast-yayp/job"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// go.uber.org/zap
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	go job.RunScheduler()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	//renderer := &TemplateRenderer{
	//	templates: template.Must(template.ParseGlob("views/*.tmpl")),
	//}
	//e.Renderer = renderer

	e.GET("/index.txt", handler.IndexTxt())
	e.GET("/api/channels", handler.GetChannels())
	e.GET("/api/channelLogs", handler.GetChannelLogs())
	e.GET("/api/channelDailyLogs", handler.GetChannelLogs())

	e.Static("/*", "public")

	e.Logger.Fatal(e.Start(":8000"))
}
