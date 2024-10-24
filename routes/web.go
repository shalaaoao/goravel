package routes

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Get("/ossUpdate", func(ctx http.Context) http.Response {

		domain := facades.Config().Env("APP_HOST", "192.168.100.15")
		port := facades.Config().Env("APP_PORT", "3001")

		url := fmt.Sprintf("http://%s:%s", domain, port)

		return ctx.Response().View().Make("ossUpdate.tmpl", map[string]any{
			"url": url,
		})
	})
}
