package auth

import "github.com/labstack/echo/v4"

type SignUpRequest struct {
}

type SignUpResponse struct {
}

// @tags			Auth
// @summary			User Signup
// @description		Creates New User Account
// @accept			json
// @produce			json
// @param			body body 		SignUpRequest true "SignUpRequest"
// @success			200	{object}	SignUpResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/v1/signup [post]
func (s *AuthService) AppleLogin(c echo.Context) error {

	return nil
}
