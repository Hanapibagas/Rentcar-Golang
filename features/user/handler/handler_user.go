package handler

import (
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"
	"StartUp-Go/utils/responses"
	"log"
	"mime/multipart"
	"net/http"
	"net/smtp"
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

func (handler *UserHandler) DeleteByUuidCostumer(c echo.Context) error {
	userId := c.Param("uuid")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid.", nil))
	}

	err := handler.userService.DeleteByUuid(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error bind data. data not valid."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (handler *UserHandler) GetByUuidCostumer(c echo.Context) error {
	userId := c.Param("uuid")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid.", nil))
	}

	result, err := handler.userService.GetByUuid(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error bind data. data not valid."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success.", result))
}

func getRoleString(role string) string {
	switch role {
	case "2":
		return "super admin"
	case "3":
		return "admin"
	case "4":
		return "driver"
	case "5":
		return "customer"
	default:
		return "owner"
	}
}

func (handler *UserHandler) GetAllCostumer(c echo.Context) error {
	result, err := handler.userService.GetAllCostumer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data."+err.Error(), nil))
	}

	var customerResponse []ListCostumerRespon
	for _, customer := range result {
		roleString := getRoleString(customer.Role)
		customerResponse = append(customerResponse, ListCostumerRespon{
			Uuid:          customer.Uuid,
			UserName:      customer.UserName,
			Status:        customer.Status,
			Role:          roleString,
			FullName:      customer.FullName,
			TempatLahir:   customer.TempatLahir,
			Alamat:        customer.Alamat,
			Email:         customer.Email,
			Notelp:        customer.Notelp,
			NotelpKerabat: customer.NotelpKerabat,
			Ktp:           customer.Ktp,
			Pekerjaan:     customer.Pekerjaan,
			FotoKtp:       customer.FotoKtp,
		})
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success.", customerResponse))
}

func (handler *UserHandler) UpdateCostumer(c echo.Context) error {
	var fileSize int64
	var nameFile string
	userId := c.Param("uuid")

	var reqData = UpdatetCostumer{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid"+errBind.Error(), nil))
	}

	costumerUpdateCore := RequestToUpdateCostumer(reqData)

	fileHeader, _ := c.FormFile("foto_ktp")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error invalid type format, format file not valid", nil))
		}

		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
		}
	}

	err := handler.userService.UpdateCustomer(userId, costumerUpdateCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Update pelanggan successful", nil))
}

func (handler *UserHandler) InsertCostumer(c echo.Context) error {
	var fileSize int64
	var nameFile string

	var reqData = InsertCostumer{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	costumerCore := RequestToInsertCostumer(reqData)

	surel_pengirim := "disdukcapilmkskota@gmail.com"
	kata_sandi := "tqozsznukogmyrdr"
	penerima := costumerCore.Email

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	emailBody := `
			<html>
			<head>
				<style>
					.container {
							font-family: Arial, sans-serif;
							max-width: 600px;
							margin: 0 auto;
							padding: 20px;
							border: 1px solid #ccc;
							border-radius: 5px;
					}
					h1 {
						text-align: center;
						color: #333;
					}
					p {
							color: #666;
					}
					.container a{
						color: white;
					}
					.button {
						display: inline-block;
						padding: 10px 20px;
						background-color: #007bff;
						color: white;
						text-decoration: none;
						border-radius: 5px;
						transition: background-color 0.3s ease;
						margin-left: 240px;
					}
						.button:hover {
						background-color: #0056b3;
						color: white;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to Our Platform</h1>
					<hr>
					<p>Hello ` + reqData.Email + `,</p>
					<p>Thank you for registering with us. We are excited to have you on board!</p>
					<p>Please verify the grow below</p>
					<a href="http://127.0.0.1:8080/verified" class="button">click me to verify</a>
				</div>
			</body>
			</html>
		`

	message := []byte("Subject: Testing Go Email\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		emailBody)

	auth := smtp.PlainAuth("", surel_pengirim, kata_sandi, smtpHost)

	fileHeader, _ := c.FormFile("foto_ktp")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error invalid type format, format file not valid", nil))
		}

		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
		}
	}

	err := handler.userService.InsertCustomer(costumerCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed", nil))
	}

	smtp.SendMail(smtpHost+":"+smtpPort, auth, surel_pengirim, []string{penerima}, message)

	return c.JSON(http.StatusOK, responses.WebResponse("Insert pelanggan successful", nil))
}

func (handler *UserHandler) InsertUser(c echo.Context) error {
	var fileSize int64
	var nameFile string

	var reqData = InsertUser{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := RequestToCoreUser(reqData)

	surel_pengirim := "disdukcapilmkskota@gmail.com"
	kata_sandi := "tqozsznukogmyrdr"
	penerima := userCore.Email

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	emailBody := `
			<html>
			<head>
				<style>
					.container {
							font-family: Arial, sans-serif;
							max-width: 600px;
							margin: 0 auto;
							padding: 20px;
							border: 1px solid #ccc;
							border-radius: 5px;
					}
					h1 {
						text-align: center;
						color: #333;
					}
					p {
							color: #666;
					}
					.container a{
						color: white;
					}
					.button {
						display: inline-block;
						padding: 10px 20px;
						background-color: #007bff;
						color: white;
						text-decoration: none;
						border-radius: 5px;
						transition: background-color 0.3s ease;
						margin-left: 240px;
					}
						.button:hover {
						background-color: #0056b3;
						color: white;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to Our Platform</h1>
					<hr>
					<p>Hello ` + reqData.Email + `,</p>
					<p>Thank you for registering with us. We are excited to have you on board!</p>
					<p>Please verify the grow below</p>
					<a href="http://127.0.0.1:8080/verified" class="button">click me to verify</a>
				</div>
			</body>
			</html>
		`

	message := []byte("Subject: Testing Go Email\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		emailBody)

	auth := smtp.PlainAuth("", surel_pengirim, kata_sandi, smtpHost)

	fileHeader, _ := c.FormFile("foto_ktp")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error invalid type format, format file not valid", nil))
		}

		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
		}
	}

	err := handler.userService.InsertUser(userCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed."+err.Error(), nil))
	}

	result := map[string]interface{}{
		"hasil":  userCore,
		"file":   file,
		"folder": nameFile,
	}

	smtp.SendMail(smtpHost+":"+smtpPort, auth, surel_pengirim, []string{penerima}, message)

	return c.JSON(http.StatusOK, responses.WebResponse("Insert user successful", result))
}

func (handler *UserHandler) VerifiedEmail(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	updateEmailVerified := UserRequestVerified{
		Status: "T",
	}

	errBind := c.Bind(&updateEmailVerified)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid"+errBind.Error(), nil))
	}

	usereCore := RequestToUpdateVerified(updateEmailVerified)

	err := handler.userService.VerifiedEmail(uint(userId), usereCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error editing data failed."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("User name verification successful", nil))
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequestRegister{}
	// log.Println("role:", newUser.)
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	user := RequestUserRegisterToCore(newUser)

	surel_pengirim := "disdukcapilmkskota@gmail.com"
	kata_sandi := "tqozsznukogmyrdr"
	penerima := user.Email

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	emailBody := `
			<html>
			<head>
				<style>
					.container {
							font-family: Arial, sans-serif;
							max-width: 600px;
							margin: 0 auto;
							padding: 20px;
							border: 1px solid #ccc;
							border-radius: 5px;
					}
					h1 {
						text-align: center;
						color: #333;
					}
					p {
							color: #666;
					}
					.container a{
						color: white;
					}
					.button {
						display: inline-block;
						padding: 10px 20px;
						background-color: #007bff;
						color: white;
						text-decoration: none;
						border-radius: 5px;
						transition: background-color 0.3s ease;
						margin-left: 240px;
					}
						.button:hover {
						background-color: #0056b3;
						color: white;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to Our Platform</h1>
					<hr>
					<p>Hello ` + newUser.Email + `,</p>
					<p>Thank you for registering with us. We are excited to have you on board!</p>
					<p>Please verify the grow below</p>
					<a href="http://127.0.0.1:8080/verified" class="button">click me to verify</a>
				</div>
			</body>
			</html>
		`

	message := []byte("Subject: Testing Go Email\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		emailBody)

	auth := smtp.PlainAuth("", surel_pengirim, kata_sandi, smtpHost)

	_, _, errRegister := handler.userService.Register(user)
	if errRegister != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+errRegister.Error(), nil))
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, surel_pengirim, []string{penerima}, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+err.Error(), nil))
	}

	responseData := UserResponRegister{
		UserName: newUser.UserName,
		Email:    newUser.Email,
		Notelp:   newUser.Notelp,
		Respon:   "Silahkan chek email anda untuk melakukan virifikasi email.",
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("register success", responseData))
}

func (handler *UserHandler) LoginUser(c echo.Context) error {
	var reqData = UserRequestLogin{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	result, token, err := handler.userService.Login(reqData.UserName, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Email atau password tidak boleh kosong "+err.Error(), nil))
	}

	log.Println("DATA:", result)

	responseData := map[string]any{
		"id":        result.ID,
		"uuid":      result.Uuid,
		"user_name": result.UserName,
		"status":    result.Status,
		"token":     token,
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Login success", responseData))
}
