package handler

import (
	"errors"
	"net/http"
	"peercast-yayp/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func GetChannelLogs() echo.HandlerFunc {
	return func(c echo.Context) error {
		cn := c.QueryParam("cn")
		date := c.QueryParam("date")

		logTime, ok := parseDateStr(date)
		if !ok {
			return errors.New("error")
		}

		db, _ := models.NewDB("root:password@/yayp?charset=utf8&parseTime=True&loc=Local")
		logs := db.FindChannelLogsByNameAndLogTime(cn, logTime)

		return c.JSON(http.StatusOK, logs)
	}
}

func parseDateStr(str string) (t time.Time, ok bool) {
	if len(str) != 8 {
		return time.Time{}, false
	}

	var ymd [3]int

	offset := 0
	for i, size := range []int{4, 2, 2} {
		v, err := strconv.Atoi(str[offset : offset+size])
		if err != nil {
			return time.Time{}, false
		}
		ymd[i] = v
		offset += size
	}

	d := time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, time.UTC)

	return d, true
}
