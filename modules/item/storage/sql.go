package storage

import "gorm.io/gorm"

type sql_store struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sql_store {
	return &sql_store{db: db}
}
