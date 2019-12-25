package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"starter-project/model"
	"starter-project/repository"
)

type userController struct {
	userRepository repository.UserRepository
}

type UserController interface {
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

func NewUserController(us repository.UserRepository) UserController {
	return &userController{us}
}

func (uc *userController) GetUser(c echo.Context)  (err error) {
	id := c.QueryParam("id")

	u, err := uc.userRepository.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(c echo.Context)  (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}
	u.ID = "1"
	uc.userRepository.Set(u.ID, u)
	return c.JSON(http.StatusOK, u)
}