package colaborator

import "albo/domain"

type Service interface {
	Get(character string) (*domain.Colaborator, error)
	Save(colaborator *domain.Colaborator) error
}
