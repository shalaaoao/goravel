package routes

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {

	domain := facades.Config().Env("APP_HOST", "192.168.100.15")
	port := facades.Config().Env("APP_PORT", "3001")
	url := fmt.Sprintf("http://%s:%s", domain, port)

	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Prefix("web").Group(func(router route.Router) {
		router.Get("/ossUpdate", func(ctx http.Context) http.Response {
			return ctx.Response().View().Make("ossUpdate.tmpl", map[string]any{
				"url": url,
			})
		})

		router.Get("/dockerComposePs", func(ctx http.Context) http.Response {
			return ctx.Response().View().Make("dockerComposePs.tmpl", map[string]any{
				"url": url,
			})
		})
	})

}
