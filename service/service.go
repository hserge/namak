package service

type ICrudService[T Email | Campaign] interface {
	All() ([]T, error)
	Get(id int) (T, error)
	Update(m *T) (T, error)
	Delete(id int) (bool, error)
	Create(m *T) (bool, error)
}
