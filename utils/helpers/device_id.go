package helpers

import (
	"github.com/labstack/echo/v4"
)

func GetDeviceID(c echo.Context) (string, bool) {
	deviceID, ok := c.Get("device_id").(string)
	return deviceID, ok
}
