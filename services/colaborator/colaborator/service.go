package colaborator

import "albo/domain"

type Service interface {
	Get(character string) (*domain.Colaborator, error)
	Save(colaborators []*domain.Colaborator) error
}
