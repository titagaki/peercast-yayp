package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"peercast-yayp/config"
	"peercast-yayp/handler"
	"peercast-yayp/job"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func setRenderer(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.tmpl")),
	}
	e.Renderer = renderer
}

func routes(e *echo.Echo) {
	e.GET("/index.txt", handler.IndexTxt())
	e.GET("/api/channels", handler.GetChannels())
	e.GET("/api/channelLogs", handler.GetChannelLogs())
	e.GET("/api/channelDailyLogs", handler.GetChannelLogs())
	e.Static("/*", "public")
}

func main() {
	configPath := flag.String("config", "config/config.toml", "path of the config file")
	flag.Parse()

	cfg, err := config.FromFile(*configPath)
	if err != nil {
		panic(err)
	}

	go job.RunScheduler()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	//setRenderer(e)
	routes(e)

	// Start server with GraceShutdown
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
			e.Logger.Info("shutting down the server.")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
