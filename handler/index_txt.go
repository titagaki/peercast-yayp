package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func IndexTxt() echo.HandlerFunc {
	return func(c echo.Context) error {
		//db, _ := models.NewDB("root:password@/yayp?charset=utf8&parseTime=True&loc=Local")
		//channels := db.FindPlayingChannels()

		var output string


		return c.String(http.StatusOK, output)
	}
}
