package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func ChannelStatistics() echo.HandlerFunc {

	params := map[string]interface{}{
		"title":  "Statistics",
	}

	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "getgmt.tmpl", params)
	}
}
