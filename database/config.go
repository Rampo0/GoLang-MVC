package database

import (
	"os"
)

type dbconf struct {
	Driver string
	User   string
	Pass   string
	Name   string
	Host   string
}

func config() dbconf {

	_db := dbconf{
		Driver: os.Getenv("DB_CONNECTION"),
		User:   os.Getenv("DB_USERNAME"),
		Pass:   os.Getenv("DB_PASSWORD"),
		Name:   os.Getenv("DB_DATABASE"),
		Host:   os.Getenv("DB_HOST"),
	}

	return _db

}
