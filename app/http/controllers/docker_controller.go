package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/http/response"
	"goravel/app/service"
)

type DockerController struct {
	//Dependent services
}

func NewDockerController() *DockerController {
	return &DockerController{
		//Inject services
	}
}

func (r *DockerController) GetDockerComposePs(ctx http.Context) http.Response {

	result := service.GetDockerComposePs()

	return response.Success(ctx, http.Json{
		"result": result,
	})
}

func (r *DockerController) Up(ctx http.Context) http.Response {

	name := ctx.Request().Input("name")

	result := service.DockerComposeUp(name)

	return response.Success(ctx, http.Json{
		"result": result,
	})
}

func (r *DockerController) Down(ctx http.Context) http.Response {

	name := ctx.Request().Input("name")

	result := service.DockerComposeDown(name)

	return response.Success(ctx, http.Json{
		"result": result,
	})
}
