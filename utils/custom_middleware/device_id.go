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
		ua_branding := c.Request().Header.Get("sec-ch-ua")
		if userAgent == "" {
			userAgent = "unknown-useragent-branding"
		}
		ua_mobile := c.Request().Header.Get("sec-ch-ua-mobile")
		if userAgent == "" {
			userAgent = "unknown-useragent-mobile"
		}

		ua_platform := c.Request().Header.Get("sec-ch-ua-platform")
		if userAgent == "" {
			userAgent = "unknown-useragent-mobile"
		}

		remoteIP := c.RealIP()

		data := userAgent + remoteIP + ua_branding + ua_mobile + ua_platform

		hash := sha256.Sum256([]byte(data))
		deviceID := hex.EncodeToString(hash[:])

		c.Set("device_id", deviceID)

		return next(c)
	}
}

func DeviceIDMiddleware() echo.MiddlewareFunc {
	return middleware
}
