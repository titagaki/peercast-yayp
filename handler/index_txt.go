package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexTxt() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "index.txt")
	}
}
