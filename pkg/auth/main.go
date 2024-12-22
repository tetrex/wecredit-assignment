package auth

import (
	"github.com/rs/zerolog"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
)

type AuthService struct {
	Mode    string
	Logger  *zerolog.Logger
	Queries *db.Queries
}

type NewAuthServiceParams struct {
	Mode    string
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewAuthService(params NewAuthServiceParams) *AuthService {
	return &AuthService{
		Mode:    params.Mode,
		Logger:  params.Logger,
		Queries: params.Queries,
	}
}
