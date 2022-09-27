package handler

import (
	"echo_crud_api/entity"
	"echo_crud_api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase usecase.IUserUseCase
}

func NewUserHandler(userUC usecase.IUserUseCase) *UserHandler {
	return &UserHandler{userUC}
}

func (userHandler UserHandler) StoreUser(c echo.Context) error {
	request := entity.UserRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	result, err := userHandler.UserUseCase.StoreUser(request)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (userHandler UserHandler) GetAllUsers(c echo.Context) error {
	result, err := userHandler.UserUseCase.GetAllUsers()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (userHandler UserHandler) GetUserById(c echo.Context) error {
	paramID, _ := strconv.Atoi(c.Param("id"))
	result, err := userHandler.UserUseCase.GetUserById(paramID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (userHandler UserHandler) UpdateUser(c echo.Context) error {
	paramID, _ := strconv.Atoi(c.Param("id"))

	request := entity.UserRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	result, err := userHandler.UserUseCase.UpdateUser(paramID, request)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (userHandler UserHandler) DeleteUser(c echo.Context) error {
	paramID, _ := strconv.Atoi(c.Param("id"))

	result, err := userHandler.UserUseCase.DeleteUser(paramID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Delete User Gagal")
	}

	return c.JSON(http.StatusOK, result)
}
