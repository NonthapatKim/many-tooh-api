package repository

import (
	"database/sql"

	"github.com/NonthapatKim/many-tooth-api/internal/core/port"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) port.Repository {
	return &repository{db: db}
}
