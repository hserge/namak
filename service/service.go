package service

import "github.com/hserge/namak/model"

type ICrudService[T model.Email | Campaign] interface {
	All() ([]T, error)
	Get(id int) (T, error)
	Update(m *T) (T, error)
	Delete(id int) (bool, error)
	Create(m *T) (bool, error)
}
