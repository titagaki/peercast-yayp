package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/repositoriy"
)

func GetChannelLogs() echo.HandlerFunc {
	return func(c echo.Context) error {
		cn := c.QueryParam("cn")
		date := c.QueryParam("date")

		logTime, ok := parseDateStr(date)
		if !ok {
			return errors.New("error")
		}

		db, err := infrastructure.NewDB(config.GetConfig())
		if err != nil {
			return err
		}
		defer db.Close()

		channelLogRepo := repositoriy.NewChannelLogRepository(db)
		logs := channelLogRepo.FindChannelLogsByNameAndLogTime(cn, logTime)

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

	d := time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, time.Local)

	return d, true
}
