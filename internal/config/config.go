package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string

	PGHost	string
	PGDB	string
	PGUSER	string
	PGPASSWORD string
	PGPORT string
	PGSSL string
}

func (c *Config) Parse() {
	godotenv.Load("infra.env")

	c.JWTSecret = os.Getenv("JWT_SECRET")
	c.PGHost = os.Getenv("POSTGRES_HOST")
	c.PGDB = os.Getenv("POSTGRES_DB")
	c.PGUSER = os.Getenv("POSTGRES_USER")
	c.PGPASSWORD = os.Getenv("POSTGRES_PASSWORD")
	c.PGPORT = os.Getenv("POSTGRES_PORT")
	c.PGSSL = os.Getenv("POSTGRES_SSL")
}