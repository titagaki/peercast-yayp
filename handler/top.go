package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func TopPage() echo.HandlerFunc {

	params := map[string]interface{}{
		"title":  "Yet Another PeerCast Yellow Pages",
	}

	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "top.tmpl", params)
	}
}
