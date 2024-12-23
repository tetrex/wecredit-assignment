// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createNewUser = `-- name: CreateNewUser :exec
INSERT INTO users (username, password, primary_device, sex, age,mobile_number)
VALUES ($1, $2, $3, $4, $5,$6)
`

type CreateNewUserParams struct {
	Username      string      `json:"username"`
	Password      string      `json:"password"`
	PrimaryDevice string      `json:"primary_device"`
	Sex           pgtype.Text `json:"sex"`
	Age           pgtype.Int4 `json:"age"`
	MobileNumber  int32       `json:"mobile_number"`
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) error {
	_, err := q.db.Exec(ctx, createNewUser,
		arg.Username,
		arg.Password,
		arg.PrimaryDevice,
		arg.Sex,
		arg.Age,
		arg.MobileNumber,
	)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, primary_device, sex, age 
FROM users 
WHERE id = $1 AND is_deleted = FALSE
`

type GetUserByIdRow struct {
	ID            int32       `json:"id"`
	Username      string      `json:"username"`
	PrimaryDevice string      `json:"primary_device"`
	Sex           pgtype.Text `json:"sex"`
	Age           pgtype.Int4 `json:"age"`
}

func (q *Queries) GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PrimaryDevice,
		&i.Sex,
		&i.Age,
	)
	return i, err
}

const getUserByMobile = `-- name: GetUserByMobile :one
SELECT id, username, mobile_number, password, primary_device, sex, age, is_deleted
FROM users
WHERE mobile_number = $1 AND is_deleted = FALSE
`

func (q *Queries) GetUserByMobile(ctx context.Context, mobileNumber int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserByMobile, mobileNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.MobileNumber,
		&i.Password,
		&i.PrimaryDevice,
		&i.Sex,
		&i.Age,
		&i.IsDeleted,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, username, mobile_number, password, primary_device, sex, age, is_deleted
FROM users
WHERE username = $1 AND is_deleted = FALSE
`

func (q *Queries) GetUserByUserName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUserName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.MobileNumber,
		&i.Password,
		&i.PrimaryDevice,
		&i.Sex,
		&i.Age,
		&i.IsDeleted,
	)
	return i, err
}

const getUserWithPassword = `-- name: GetUserWithPassword :one
SELECT id, username, mobile_number, password, primary_device, sex, age, is_deleted
FROM users
WHERE username = $1
  AND password = $2
  AND is_deleted = FALSE
  AND mobile_number = $3
`

type GetUserWithPasswordParams struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	MobileNumber int32  `json:"mobile_number"`
}

func (q *Queries) GetUserWithPassword(ctx context.Context, arg GetUserWithPasswordParams) (User, error) {
	row := q.db.QueryRow(ctx, getUserWithPassword, arg.Username, arg.Password, arg.MobileNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.MobileNumber,
		&i.Password,
		&i.PrimaryDevice,
		&i.Sex,
		&i.Age,
		&i.IsDeleted,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, username, primary_device, sex, age 
FROM users 
WHERE is_deleted = FALSE
ORDER BY id ASC
LIMIT $1 OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersRow struct {
	ID            int32       `json:"id"`
	Username      string      `json:"username"`
	PrimaryDevice string      `json:"primary_device"`
	Sex           pgtype.Text `json:"sex"`
	Age           pgtype.Int4 `json:"age"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersRow{}
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.PrimaryDevice,
			&i.Sex,
			&i.Age,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isValidMobile = `-- name: IsValidMobile :one
SELECT COUNT(*) > 0 AS is_taken
FROM users
WHERE mobile_number = $1 AND is_deleted = FALSE
`

func (q *Queries) IsValidMobile(ctx context.Context, mobileNumber int32) (bool, error) {
	row := q.db.QueryRow(ctx, isValidMobile, mobileNumber)
	var is_taken bool
	err := row.Scan(&is_taken)
	return is_taken, err
}

const softDelete = `-- name: SoftDelete :exec
UPDATE users 
SET is_deleted = TRUE 
WHERE id = $1
`

func (q *Queries) SoftDelete(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, softDelete, id)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users 
SET username = $1, age = $2 
WHERE id = $3 AND is_deleted = FALSE
`

type UpdateUserParams struct {
	Username string      `json:"username"`
	Age      pgtype.Int4 `json:"age"`
	ID       int32       `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.Username, arg.Age, arg.ID)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users 
SET password = $1 
WHERE id = $2 AND is_deleted = FALSE
`

type UpdateUserPasswordParams struct {
	Password string `json:"password"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.Password, arg.ID)
	return err
}

const userNameTaken = `-- name: UserNameTaken :one
SELECT COUNT(*) > 0 AS is_taken
FROM users
WHERE username = $1 AND is_deleted = FALSE
`

func (q *Queries) UserNameTaken(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRow(ctx, userNameTaken, username)
	var is_taken bool
	err := row.Scan(&is_taken)
	return is_taken, err
}
