package colaborator

import "albo/domain"

type Repository interface {
	Get(character string) (*domain.Colaborator, error)
	Save(data *domain.Colaborator) error
}
