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

type ResendOtpRequest struct {
	UserName string `json:"user_name" validate:"required"`
}

// @tags			Auth
// @summary			Re-Send Otp
// @description		Re-Sends Otp To User
// @accept			json
// @produce			json
// @param			body body 		ResendOtpRequest true "ResendOtpRequest"
// @success			200	{object}	response.SuccessResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/v1/resend-otp [post]
func (s *AuthService) ResendOtp(c echo.Context) error {
	var req ResendOtpRequest
	if err := validate.BindAndValidate(c, &req); err != nil {
		s.Logger.Error().Err(err).Msg("Request handling failed")
		response := response.ErrResp(fmt.Sprintf("Failed to bind request = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	// check is username exists
	isUserNameValid, err := s.Queries.UserNameTaken(c.Request().Context(), req.UserName)
	if err != nil {
		s.Logger.Error().Err(err).Msg("worng username")
		response := response.ErrResp(fmt.Sprintf("worng username = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	if !isUserNameValid {
		s.Logger.Error().Err(err).Msg("worng username")
		response := response.ErrResp("worng username")
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := s.Queries.GetUserByUserName(c.Request().Context(), req.UserName)
	if err != nil {
		s.Logger.Error().Err(err).Msg("user not found")
		response := response.ErrResp(fmt.Sprintf("user not found = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	// check otp
	vaild_otp, err := s.Queries.IsValidOtp(c.Request().Context(), req.UserName)
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to get otp")
		response := response.ErrResp(fmt.Sprintf("failed to get otp = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}
	if vaild_otp {
		// there is a valid non expired otp in db , try after 5 minutes
		// dont send new
		s.Logger.Error().Err(err).Msg("there is a valid non expired otp in db , try after 5 minutes")
		response := response.ErrResp("there is a valid non expired otp in db , try after 5 minutes")
		return c.JSON(http.StatusOK, response)
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
