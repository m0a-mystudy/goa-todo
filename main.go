//go:generate goagen bootstrap -d github.com/m0a-mystudy/goa-todo/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/m0a-mystudy/goa-todo/app"
	"github.com/m0a-mystudy/goa-todo/controllers"
	"github.com/m0a-mystudy/goa-todo/store"
)

func main() {
	// Create service
	service := goa.New("Todo API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "todo" controller
	c := controllers.NewTodoController(service, store.NewDB())
	app.MountTodoController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
