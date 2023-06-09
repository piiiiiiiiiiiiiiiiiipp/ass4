package use_case

import (
	"awesomeProject3/services/contact/internal/domain/contact"
	"awesomeProject3/services/contact/internal/use-case/adapters/repository"
)

type groupUseCase struct {
	groupRepo repository.GroupRepository
	db        repository.DBRepository
}

func (g2 *groupUseCase) CreateGroup(g *contact.Group) (*contact.Group, error) {
	g, err := g2.groupRepo.Create(g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g2 *groupUseCase) ReadGroup(ID int) (*contact.Group, error) {
	g, err := g2.groupRepo.Read(ID)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g2 *groupUseCase) AddContactToGroup(telephone string, id int) error {
	err := g2.groupRepo.AddContactToGroup(telephone, id)
	if err != nil {
		return err
	}
	return nil
}

type Group interface {
	CreateGroup(g *contact.Group) (*contact.Group, error)
	ReadGroup(ID int) (*contact.Group, error)
	AddContactToGroup(telephone string, id int) error
}

func NewGroupUseCase(g repository.GroupRepository, db repository.DBRepository) Group {
	return &groupUseCase{g, db}
}
