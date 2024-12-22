package custommiddleware

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/labstack/echo/v4"
)

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userAgent := c.Request().Header.Get("User-Agent")
		if userAgent == "" {
			userAgent = "unknown"
		}
		remoteIP := c.RealIP()

		data := userAgent + remoteIP

		hash := sha256.Sum256([]byte(data))
		deviceID := hex.EncodeToString(hash[:])

		c.Set("device_id", deviceID)

		return next(c)
	}
}

func DeviceIDMiddleware() echo.MiddlewareFunc {
	return middleware
}
