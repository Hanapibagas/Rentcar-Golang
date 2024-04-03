package handler

import (
	"StartUp-Go/app/middlewares"
	"StartUp-Go/features/user"
	"StartUp-Go/features/user/data"
	"StartUp-Go/utils/responses"
	"log"
	"net/http"
	"net/smtp"

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

func convertDataUserToUserCore(input data.User) user.UserCore {
	return user.UserCore{
		EmailVerification: input.EmailVerification,
	}
}

func (handler *UserHandler) VerifiedEmail(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	var input data.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+err.Error(), nil))
	}

	userCoreInput := convertDataUserToUserCore(input)

	err := handler.userService.VerifiedEmail(userId, userCoreInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("insert success", nil))
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequestRegister{}
	// log.Println("role:", newUser.Name)
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	user := RequestUserRegisterToCore(newUser)
	log.Println("email:", user.Email)

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
					<p>Hello ` + newUser.Name + `,</p>
					<p>Thank you for registering with us. We are excited to have you on board!</p>
					<p>Please verify the grow below</p>
					<a href="https://www.example.com" class="button">click me to verify</a>
				</div>
			</body>
			</html>
		`

	message := []byte("Subject: Testing Go Email\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		emailBody)

	auth := smtp.PlainAuth("", surel_pengirim, kata_sandi, smtpHost)

	_, token, errRegister := handler.userService.Register(user)
	if errRegister != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+errRegister.Error(), nil))
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, surel_pengirim, []string{penerima}, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+err.Error(), nil))
	}

	responseData := UserResponRegister{
		Name:       newUser.Name,
		Occupation: newUser.Occupation,
		Email:      newUser.Email,
		Role:       user.Role,
		Token:      token,
		Respon:     "Silahkan chek email anda untuk melakukan virifikasi email.",
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("insert success", responseData))
}

func (handler *UserHandler) LoginUser(c echo.Context) error {
	var reqData = UserRequestLogin{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Email atau password tidak boleh kosong "+err.Error(), nil))
	}

	responData := map[string]any{
		"id":                 result.ID,
		"name":               result.Name,
		"email":              result.Email,
		"role":               result.Role,
		"email_verification": result.EmailVerification,
		"toke":               token,
	}

	return c.JSON(http.StatusOK, responses.WebResponse("insert success", responData))
}
