package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	"github.com/tetrex/wecredit-assignment/utils/otp"
	"github.com/tetrex/wecredit-assignment/utils/response"
	"github.com/tetrex/wecredit-assignment/utils/validate"
)

type OtpRequest struct {
	MobileNumber int `json:"mobile_number" validate:"required"`
}

// @tags			Auth
// @summary			Send Otp
// @description		Sends Otp To User
// @accept			json
// @produce			json
// @param			body body 		OtpRequest true "OtpRequest"
// @success			200	{object}	response.SuccessResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/v1/generate-otp [post]
func (s *AuthService) GenerateOtp(c echo.Context) error {
	var req OtpRequest
	if err := validate.BindAndValidate(c, &req); err != nil {
		s.Logger.Error().Err(err).Msg("Request handling failed")
		response := response.ErrResp(fmt.Sprintf("Failed to bind request = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	// check is username exists
	isValidMobileNumber, err := s.Queries.IsValidMobile(c.Request().Context(), int32(req.MobileNumber))
	if err != nil {
		s.Logger.Error().Err(err).Msg("worng mobile number")
		response := response.ErrResp(fmt.Sprintf("worng mobile number = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	if !isValidMobileNumber {
		s.Logger.Error().Err(err).Msg("worng mobile number")
		response := response.ErrResp("worng mobile number")
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := s.Queries.GetUserByMobile(c.Request().Context(), int32(req.MobileNumber))
	if err != nil {
		s.Logger.Error().Err(err).Msg("user not found")
		response := response.ErrResp(fmt.Sprintf("user not found = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}
	new_otp, _ := otp.NewOtp(6)
	// send Otp
	err = s.Queries.CreateNewOtp(c.Request().Context(), db.CreateNewOtpParams{
		UserID: user.ID,
		Otp:    new_otp,
	})
	if err != nil {
		s.Logger.Error().Err(err).Msg("otp creation failed")
		response := response.ErrResp(fmt.Sprintf("otp creation failed = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response.OkResp("ok", response.OkResp("ok, otp sent", struct{}{})))
}
