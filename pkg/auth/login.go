package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	"github.com/tetrex/wecredit-assignment/utils/helpers"
	"github.com/tetrex/wecredit-assignment/utils/jwt"
	"github.com/tetrex/wecredit-assignment/utils/password"
	"github.com/tetrex/wecredit-assignment/utils/response"
	"github.com/tetrex/wecredit-assignment/utils/validate"
)

type LoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Otp      string `json:"otp" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	NewDevice   bool   `json:"is_new_device"`
	OldDeviceId string `json:"old_device_id"`
	NewDeviceId string `json:"new_device_id"`
}

// @tags			Auth
// @summary			User Login
// @description		Login for User
// @accept			json
// @produce			json
// @param			body body 		LoginRequest true "LoginRequest"
// @success			200	{object}	LoginResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/v1/login [post]
func (s *AuthService) Login(c echo.Context) error {
	var req LoginRequest
	if err := validate.BindAndValidate(c, &req); err != nil {
		s.Logger.Error().Err(err).Msg("Request handling failed")
		response := response.ErrResp(fmt.Sprintf("Failed to bind request = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	//
	hashed_password := password.HashPassword(req.Password)

	user, err := s.Queries.GetUserWithPassword(c.Request().Context(), db.GetUserWithPasswordParams{
		Username: req.UserName,
		Password: hashed_password,
	})
	if err != nil {
		s.Logger.Error().Err(err).Msg("worng username or password")
		response := response.ErrResp(fmt.Sprintf("worng username or password = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	// check otp
	isOtpValid, err := s.Queries.CheckOtp(c.Request().Context(), db.CheckOtpParams{
		UserID: user.ID,
		Otp:    req.Otp,
	})
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to get otp")
		response := response.ErrResp(fmt.Sprintf("failed to get otp = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}
	if !isOtpValid {
		response := response.ErrResp("wrong Otp")
		return c.JSON(http.StatusBadRequest, response)
	}

	err = s.Queries.MarkOtpUsed(c.Request().Context(), db.MarkOtpUsedParams{
		Otp:      req.Otp,
		Username: req.UserName,
	})
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to mark otp used")
		response := response.ErrResp(fmt.Sprintf("failed to mark otp used = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	// all ok from here
	old_device_id := user.PrimaryDevice
	new_device_id, _ := helpers.GetDeviceID(c)
	isSameDevice := old_device_id == new_device_id

	token, err := jwt.GenerateTokens(int(user.ID))
	if err != nil {
		s.Logger.Error().Err(err).Msg("error genetaring jwt token")
		response := response.ErrResp(fmt.Sprintf("error genetaring jwt token = %s", err.Error()))
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response.OkResp("ok", LoginResponse{AccessToken: token.AccessCode, NewDevice: !isSameDevice, OldDeviceId: old_device_id, NewDeviceId: new_device_id}))
}
