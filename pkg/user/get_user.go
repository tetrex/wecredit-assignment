package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	"github.com/tetrex/wecredit-assignment/utils/jwt"
	"github.com/tetrex/wecredit-assignment/utils/response"
)

type User struct {
	ID            int32  `json:"id"`
	Username      string `json:"username"`
	PrimaryDevice string `json:"primary_device"`
	Sex           string `json:"sex"`
	Age           int32  `json:"age"`
}
type UserResponse struct {
	Msg       string `json:"msg"`
	User_Data User   `json:"user_data"`
}

func responseFmt(data db.GetUserByIdRow) UserResponse {
	return UserResponse{
		Msg: "ok",
		User_Data: User{
			ID:            data.ID,
			Username:      data.Username,
			PrimaryDevice: data.PrimaryDevice,
			Sex:           data.Sex.String,
			Age:           data.Age.Int32,
		},
	}
}

// @tags			User
// @summary			User Deatils
// @description		Get User detail by id
// @accept			json
// @produce			json
// @success			200	{object}	UserResponse
// @failure			500	{object}	response.ErrorResponse
// @router			/v1/user [get]
func (s *UserService) GetUserById(c echo.Context) error {
	user, ok := jwt.GetJwt(c)
	if !ok {
		s.Logger.Error().Err(fmt.Errorf("jwt error")).Msg("Failed to get jwt")
		response := response.ErrResp(fmt.Errorf("jwt error"))
		return c.JSON(http.StatusUnauthorized, response)
	}

	user_data, err := s.Queries.GetUserById(c.Request().Context(), int32(user.UserId))
	if err != nil {
		s.Logger.Error().Err(err).Msg("no user found")
		response := response.ErrResp(fmt.Sprintf("no user found = %s", err.Error()))
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response.OkResp("ok", responseFmt(user_data)))
}
