package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers"
)

func Api() {

	facades.Route().Prefix("api").Group(func(router route.Router) {

		userController := controllers.NewUserController()
		router.Get("/users/{id}", userController.Show)

		ossController := controllers.NewOssController()
		router.Get("/oss/show", ossController.Show)
		router.Post("/oss/update", ossController.Update)

		testController := controllers.NewTestController()
		router.Get("test", testController.Test)

		dockerController := controllers.NewDockerController()
		router.Get("docker/getDockerComposePs", dockerController.GetDockerComposePs)
		router.Post("docker/up", dockerController.Up)
		router.Post("docker/down", dockerController.Down)
	})

}
