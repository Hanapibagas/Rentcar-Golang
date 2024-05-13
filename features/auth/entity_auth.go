package auth

type RegisterCore struct {
	ID       uint
	UserName string
	Password string
	Status   string
	Role     int
	FullName string
	Email    string
	Notelpn  string
}

type LoginCore struct {
	ID       uint
	UserName string
	Password string
	Uuid     string
	Status   string
}

type AuthDataInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(username, password string) (data *LoginCore, err error)
	VerifiedEmail(id uint, input RegisterCore) error
}

type AuthServiceInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(username, password string) (data *LoginCore, token string, err error)
	VerifiedEmail(id uint, input RegisterCore) error
}
