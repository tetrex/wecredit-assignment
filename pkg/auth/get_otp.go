package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tetrex/wecredit-assignment/utils/response"
	"github.com/tetrex/wecredit-assignment/utils/validate"
)

// @tags			Auth
// @summary			Get Otp
// @description		Gets Otp To User
// @accept			json
// @produce			json
// @param			body body 		OtpRequest true "OtpRequest"
// @success			200	{object}	response.SuccessResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/v1/get-valid-otp [post]
func (s *AuthService) GetOtp(c echo.Context) error {
	var req OtpRequest
	if err := validate.BindAndValidate(c, &req); err != nil {
		s.Logger.Error().Err(err).Msg("Request handling failed")
		response := response.ErrResp(fmt.Sprintf("Failed to bind request = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}
	data, err := s.Queries.GetValidOtpForUserName(c.Request().Context(), req.UserName)
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to get otp")
		response := response.ErrResp(fmt.Sprintf("failed to get otp = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response.OkResp("ok", response.OkResp("ok, otp sent", data)))
}
