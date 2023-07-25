package store

import (
	"github.com/jmoiron/sqlx"
)

type UserStore struct{
	db sqlx.DB
}