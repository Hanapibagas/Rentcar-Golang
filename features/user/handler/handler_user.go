package handler

import (
	"StartUp-Go/features/user"
	"StartUp-Go/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func NewUser(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequestRegister{}
	// log.Println("role:", newUser.Name)
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	user := RequestUserRegisterToCore(newUser)
	_, token, errRegister := handler.userService.Register(user)
	if errRegister != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+errRegister.Error(), nil))
	}
	responseData := UserResponRegister{
		Name:       newUser.Name,
		Occupation: newUser.Occupation,
		Email:      newUser.Email,
		Role:       user.Role,
		Token:      token,
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("insert success", responseData))
}
