package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// GetChannelLogs は指定チャンネル・日付のログを返すAPIハンドラー。
func (h *Handler) GetChannelLogs(c echo.Context) error {
	name := c.QueryParam("cn")
	dateStr := c.QueryParam("date")

	date, err := parseDate(dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid date format, expected YYYYMMDD")
	}

	logs, err := h.deps.ChannelLogs.FindByNameAndDate(c.Request().Context(), name, date)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch channel logs")
	}

	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	return c.JSON(http.StatusOK, logs)
}

func parseDate(s string) (time.Time, error) {
	if len(s) != 8 {
		return time.Time{}, fmt.Errorf("date must be 8 characters")
	}

	var ymd [3]int
	offset := 0
	for i, size := range []int{4, 2, 2} {
		v, err := strconv.Atoi(s[offset : offset+size])
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid date component: %w", err)
		}
		ymd[i] = v
		offset += size
	}

	return time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, time.Local), nil
}
