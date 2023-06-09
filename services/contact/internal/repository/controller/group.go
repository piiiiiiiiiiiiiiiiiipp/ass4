package controller

import (
	use_case "architecture_go/services/contact/internal/use-case"
)

type groupController struct {
	groupUseCase use_case.Group
}

func (gg groupController) CreateGroup(ctx Context) {
	//TODO implement me
	panic("implement me")
}

func (gg groupController) ReadGroup(ctx Context) {
	//TODO implement me
	panic("implement me")
}

func (gg groupController) AddContactToGroup(ctx Context) {
	//TODO implement me
	panic("implement me")
}

type Group interface {
	CreateGroup(ctx Context)
	ReadGroup(ctx Context)
	AddContactToGroup(ctx Context)
}

func NewGroupsController(cs use_case.Group) Group {
	return &groupController{cs}
}
