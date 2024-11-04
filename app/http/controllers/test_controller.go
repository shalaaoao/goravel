package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type TestController struct {
	//Dependent services
}

func NewTestController() *TestController {
	return &TestController{
		//Inject services
	}
}

func (r *TestController) Test(ctx http.Context) http.Response {

	result, _ := facades.Storage().AllFiles("./logs")

	return ctx.Response().Success().Json(http.Json{
		"result": result,
	})
}
