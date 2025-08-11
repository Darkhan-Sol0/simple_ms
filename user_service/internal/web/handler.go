package web

import (
	"myApp/templates/temp_components"

	"github.com/labstack/echo/v4"
)

// type (
// 	Responce struct {
// 		Data any `json:"data"`
// 	}
// )

func (r *routingConfig) indexHandler(ctx echo.Context) error {
	responce := r.service.Hello()
	templ := temp_components.TextOut(responce)
	return templ.Render(ctx.Request().Context(), ctx.Response())
}

func (r *routingConfig) listHandler(ctx echo.Context) error {
	responce := r.service.GetList()
	templ := temp_components.ListOut(responce)
	return templ.Render(ctx.Request().Context(), ctx.Response())
}
