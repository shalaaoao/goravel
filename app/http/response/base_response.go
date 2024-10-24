package response

import (
	"github.com/goravel/framework/contracts/http"
)

func Success(ctx http.Context, data interface{}) http.Response {

	return ctx.Response().Success().Json(http.Json{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func Fail(ctx http.Context, code int, message string) http.Response {

	return ctx.Response().Success().Json(http.Json{
		"code": 200,
		"msg":  "success",
		"data": nil,
	})
}
