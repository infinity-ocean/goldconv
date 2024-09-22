package repo

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakePool() (*pgxpool.Pool, error) {
	// const DB = "postgres://postgres:12345@localhost:5432/postgres"
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		fmt.Println("Unable to parse config: ", err)
	}
	
	config.ConnConfig.Host = os.Getenv("POSTGRES_HOST")
	portStr := os.Getenv("POSTGRES_PORT")
	portInt, _ := strconv.Atoi(portStr)
	config.ConnConfig.Port = uint16(portInt) 
	config.ConnConfig.User = os.Getenv("POSTGRES_USER")
	config.ConnConfig.Password = os.Getenv("POSTGRES_PASSWORD")
	config.ConnConfig.Database = os.Getenv("POSTGRES_DATABASE")
	config.ConnConfig.RuntimeParams["sslmode"] = os.Getenv("POSTGRES_SSL")

	dbpool, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
	    fmt.Println("Unable to create connection pool: "+err.Error())
		return nil, err
	}
	return dbpool, nil
}
