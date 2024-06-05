package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUIntParam(c echo.Context, key string) (uint, error) {
	id := c.Param(key)

	uid64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(uid64), nil
}
