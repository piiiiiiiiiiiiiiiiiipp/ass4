package postgres

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/use-case/adapters/repository"
	"github.com/jinzhu/gorm"
)

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
	return &groupRepository{db}
}

func (gg groupRepository) Create(g *contact.Group) (*contact.Group, error) {
	err := gg.db.Create(&g).Error
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (gg groupRepository) Read(ID int) (*contact.Group, error) {
	err := gg.db.Find(&ID).Error
	if err != nil {
		return nil, err
	}
	//nuzno izmenit
	return nil, err
}

func (gg groupRepository) AddContactToGroup(telephone string, id int) error {
	//TODO implement me
	panic("implement AddContactToGroup how?")
}
