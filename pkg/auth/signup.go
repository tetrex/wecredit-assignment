package auth

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	"github.com/tetrex/wecredit-assignment/utils/helpers"
	"github.com/tetrex/wecredit-assignment/utils/password"
	"github.com/tetrex/wecredit-assignment/utils/response"
	"github.com/tetrex/wecredit-assignment/utils/validate"
)

type SignUpRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
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
func (s *AuthService) SignUp(c echo.Context) error {
	var req SignUpRequest
	if err := validate.BindAndValidate(c, &req); err != nil {
		s.Logger.Error().Err(err).Msg("Request handling failed")
		response := response.ErrResp(fmt.Sprintf("Failed to bind request = %s", err.Error()))
		return c.JSON(http.StatusBadRequest, response)
	}

	// check if username is taken or not
	isTaken, err := s.Queries.UserNameTaken(c.Request().Context(), req.UserName)
	if err != nil {
		s.Logger.Error().Err(err).Msg("username check failed")
		response := response.ErrResp(fmt.Sprintf("username check failed = %s", err.Error()))
		return c.JSON(http.StatusInternalServerError, response)
	}

	if isTaken {
		return c.JSON(http.StatusOK, response.SuccessResponse{Data: "error/username_taken", Msg: fmt.Sprintf("user name %s , is taken", req.UserName)})
	}

	device_id, _ := helpers.GetDeviceID(c)

	// create user
	err = s.Queries.CreateNewUser(c.Request().Context(), db.CreateNewUserParams{
		Username:      req.UserName,
		Password:      password.HashPassword(req.Password),
		PrimaryDevice: device_id,
		Sex:           pgtype.Text{String: req.Sex, Valid: true},
		Age:           pgtype.Int4{Int32: req.Age, Valid: true},
	})
	if err != nil {
		s.Logger.Error().Err(err).Msg("new user creation failed")
		response := response.ErrResp(fmt.Sprintf("new user creation failed = %s", err.Error()))
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Data: "success/user-created", Msg: "user created"})
}
