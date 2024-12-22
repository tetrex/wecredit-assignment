package main

import (
	"context"
	"errors"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import file source for migrations
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	"github.com/tetrex/wecredit-assignment/utils/config"
	"github.com/tetrex/wecredit-assignment/utils/logger"
)

// @title			server api
// @version			1.0
// @description		This is a backend api server
// @contact.name	github.com/tetrex
// @license.name	MIT License
//
// @host			localhost:8000
// @basePath		/
func main() {
	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config ")
		log.Fatal(err.Error())
	}

	// logger
	l := logger.New(config.AppEnv)

	// pg connection
	pg_config, err := pgxpool.ParseConfig(config.PgConnStr)
	if err != nil {
		l.Fatal().Err(errors.New("cannot connect to to db"))
	}

	pg_config.MaxConns = 20                     // Maximum number of connections in the pool
	pg_config.MaxConnLifetime = 5 * time.Minute // Maximum lifetime of a connection
	pg_config.MaxConnIdleTime = 2 * time.Minute // Maximum time a connection can remain idle

	// pg quries and pool
	db_pool, err := pgxpool.NewWithConfig(context.Background(), pg_config)
	if err != nil {
		l.Fatal().Err(errors.New("cannot connect to db db_pool"))
	}
	defer db_pool.Close()
	pg_queries := db.New(db_pool)

	// new server
}
