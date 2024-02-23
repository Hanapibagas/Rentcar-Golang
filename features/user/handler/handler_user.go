package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/user"
	"BE-REPO-20/utils/responses"
	"mime/multipart"
	"net/http"
	"strings"

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

func (handler *UserHandler) SelectUser(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	data, err := handler.userService.SelectUser(idJWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	result := CoreToResponse(*data)

	return c.JSON(http.StatusOK, responses.WebResponse("read success", result))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	var fileSize int64
	var nameFile string
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	var reqData = UserRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind, data not valid", nil))
	}
	userCore := RequestToCore(reqData)

	// upload img
	fileHeader, _ := c.FormFile("image_url")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		// get data type
		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		// data type validation
		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error invalid type format, format file not valid", nil))
		}

		// data size validation max 2mb
		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
		}
	}

	err := handler.userService.UpdateUser(idJWT, userCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) UpdateShop(c echo.Context) error {
	var fileSize int64
	var nameFile string
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	var reqData = UserShopRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind, data not valid", nil))
	}
	userCore := RequestToCoreShop(reqData)

	// upload img
	fileHeader, _ := c.FormFile("image_url")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		// get data type
		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		// data type validation
		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error invalid type format, format file not valid", nil))
		}

		// data size validation max 2mb
		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
		}
	}
	err := handler.userService.UpdateShop(idJWT, userCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) SelectShop(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}
	data, err := handler.userService.SelectShop(idJWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	result := CoreToResponseShop(*data)

	return c.JSON(http.StatusOK, responses.WebResponse("read success", result))
}

func (handler *UserHandler) Delete(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}
	err := handler.userService.Delete(idJWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. delete failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
