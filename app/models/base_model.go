package models

import "github.com/jackc/pgx/v5"

type CrudRepository[T any] interface {
	All() ([]T, error)
	Get(id int) (T, error)
	Update(m *T) (T, error)
	Delete(id int) (bool, error)
	Create(m *T) (bool, error)
}

type Repository[T CrudRepository[T]] struct {
	db *pgx.Conn
}

func NewRepository[T CrudRepository[T]](db *pgx.Conn) *Repository[T] {
	return &Repository[T]{
		db,
	}
}
