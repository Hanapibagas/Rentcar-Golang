package user

type UserCore struct {
	ID                uint `json:"id"`
	Name              string
	Occupation        string
	Email             string `gorm:"default:null;unique"`
	EmailVerification string
	AvatarFileName    string
	Password          string
	Role              string
}

type EmailVerification struct {
	ID                uint
	EmailVerification string
}

type UserDataInterface interface {
	Register(input UserCore) (data *UserCore, token string, err error)
	Login(email, password string) (data *UserCore, err error)
	VerifiedEmail(id uint, input EmailVerification) error
	UpdatePassword(id uint, input UserCore) error
	CheckPassword(savedPassword, inputPassword string) bool
}

type UserServiceInterface interface {
	Register(input UserCore) (data *UserCore, token string, err error)
	Login(email, password string) (data *UserCore, token string, err error)
	VerifiedEmail(id uint, input EmailVerification) error
	UptdatePassword(id uint, input UserCore) error
}
