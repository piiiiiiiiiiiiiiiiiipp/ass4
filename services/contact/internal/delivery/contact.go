package delivery

import (
	"awesomeProject3/services/contact/internal/repository/controller"
	"awesomeProject3/services/contact/internal/repository/storage/postgres"
	use_case "awesomeProject3/services/contact/internal/use-case"
)

func (r *registry) NewContactController() controller.Contact {
	c := use_case.NewContactUseCase(
		postgres.NewContactRepository(r.db),
		postgres.NewDBRepository(r.db),
	)
	return controller.NewContactsController(c)
}
