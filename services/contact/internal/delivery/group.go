package delivery

import (
	"awesomeProject3/services/contact/internal/repository/controller"
	"awesomeProject3/services/contact/internal/repository/storage/postgres"
	use_case "awesomeProject3/services/contact/internal/use-case"
)

func (r *registry) NewGroupController() controller.Group {
	g := use_case.NewGroupUseCase(
		postgres.NewGroupRepository(r.db),
		postgres.NewDBRepository(r.db),
	)
	return controller.NewGroupsController(g)
}
