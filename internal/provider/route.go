package provider

import (
	"star/internal"
	"star/routes"
)

func Route(app *internal.Application) {
	route := app.GetRoute()
	route.Use()
	{
		api := route.Group("/api")
		{
			routes.ApiV1(api)
		}
	}
}
