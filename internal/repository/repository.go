package repository

import (
	"github.com/go-pg/pg"
)

// Repository is
type Repository struct {
	db *pg.DB
}

// NewRepository is
func NewRepository(db *pg.DB) (IRepository, error) {
	return &Repository{db: db}, nil
}
