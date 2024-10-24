package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	ossController := controllers.NewOssController()
	facades.Route().Get("/oss/show", ossController.Show)
	facades.Route().Post("/oss/update", ossController.Update)
}
