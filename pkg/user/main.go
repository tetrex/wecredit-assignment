package user

import (
	"github.com/rs/zerolog"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
)

type UserService struct {
	Mode    string
	Logger  *zerolog.Logger
	Queries *db.Queries
}

type NewUserServiceParams struct {
	Mode    string
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewUserService(params NewUserServiceParams) *UserService {
	return &UserService{
		Mode:    params.Mode,
		Logger:  params.Logger,
		Queries: params.Queries,
	}
}
