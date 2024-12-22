package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tetrex/wecredit-assignment/utils/response"
)

type HealthService struct{}

func NewHealthService() *HealthService {

	return &HealthService{}
}

// @tags			Health
// @summary			Get Health check status
// @description		returns server time
// @accept			json
// @produce			json
// @success			200	{object}	int64
// @failure			500	{object}	utils.ErrorResponse
// @router			/ [get]
func (s *HealthService) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, response.OkResp("all ok, server time", struct{ Data int64 }{Data: time.Now().UnixNano()}))
}
