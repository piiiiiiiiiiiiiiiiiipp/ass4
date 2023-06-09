package router

import (
	"architecture_go/services/contact/internal/repository/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/contacts", func(context echo.Context) error {
		return c.Contact.GetContacts(context)
	})
	e.POST("/contacts", func(context echo.Context) error {
		return c.Contact.CreateContacts(context)
	})
	e.DELETE("/contacts", func(context echo.Context) error {
		return c.Contact.DeleteContacts(context)
	})
	e.POST("/contacts", func(context echo.Context) error {
		return c.Contact.UpdateContacts(context)
	})
	e.GET("/groups", func(context echo.Context) error {
		// not yet implement
		//return c.Group.ReadGroup(context)
		return nil
	})
	e.POST("/groups", func(context echo.Context) error {
		//not yet implement
		//return c.Group.CreateGroup(context)
		return nil
	})
	e.POST("/groups", func(context echo.Context) error {
		//not yet implement
		//return c.Group.AddContactToGroup(context)
		return nil
	})
	return e
}
