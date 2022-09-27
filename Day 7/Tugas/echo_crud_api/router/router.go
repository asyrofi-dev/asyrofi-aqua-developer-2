package router

import (
	"echo_crud_api/handler"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, handler *handler.UserHandler) {
	e.POST("/", handler.StoreUser)
	e.GET("/", handler.GetAllUsers)
	e.GET("/:id", handler.GetUserById)
	e.PUT("/:id", handler.UpdateUser)
	e.DELETE("/:id", handler.DeleteUser)
}
