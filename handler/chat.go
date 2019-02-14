package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Chat() echo.HandlerFunc {

	params := map[string]interface{}{
		"title":  "Chat",
	}

	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "chat.tmpl", params)
	}
}