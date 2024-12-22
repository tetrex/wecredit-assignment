package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tetrex/wecredit-assignment/utils/helpers"
	"github.com/tetrex/wecredit-assignment/utils/response"
)

type HealthService struct{}

func NewHealthService() *HealthService {

	return &HealthService{}
}

type HealthResponse struct {
	ServerTime int64  `json:"server_time"` // Server time in Unix timestamp (seconds)
	DeviceId   string `json:"device_id"`   // Server time in Unix timestamp (seconds)
	Msg        string `json:"msg"`         // Message
}

// @tags			Health
// @summary			Get Health check status
// @description		returns server time
// @accept			json
// @produce			json
// @success			200	{object}	HealthResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/ [get]
func (s *HealthService) HealthCheck(c echo.Context) error {
	device_id, _ := helpers.GetDeviceID(c)
	return c.JSON(http.StatusOK, response.OkResp("all ok, server time", HealthResponse{Msg: "server unix time (sec)", ServerTime: time.Now().Unix(), DeviceId: device_id}))
}
