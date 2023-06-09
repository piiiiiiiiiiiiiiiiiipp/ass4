package delivery

import (
	"awesomeProject3/services/contact/internal/repository/controller"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

func (r registry) NewAppController() controller.AppController {
	return controller.AppController{
		Contact: r.NewContactController(),
		Group:   r.NewGroupController(),
	}
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}
