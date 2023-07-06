package service

import (
	"albo/colaborator"
	"albo/domain"
)

type colaboratorService struct {
	Repo colaborator.Repository
}

func NewColaboratorService(repo colaborator.Repository) *colaboratorService {
	return &colaboratorService{
		Repo: repo,
	}
}

func (srv *colaboratorService) Get(character string) (*domain.Colaborator, error) {
	return srv.Repo.Get(character)
}

func (srv *colaboratorService) Save(data []*domain.Colaborator) error {
	return srv.Repo.Save(data)
}
