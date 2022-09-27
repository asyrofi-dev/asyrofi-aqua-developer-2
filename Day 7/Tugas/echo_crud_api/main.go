package main

import (
	"echo_crud_api/config"
	"echo_crud_api/handler"
	"echo_crud_api/repository"
	"echo_crud_api/router"
	"echo_crud_api/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.Migrate()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	e := echo.New()

	router.Router(e, userHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
