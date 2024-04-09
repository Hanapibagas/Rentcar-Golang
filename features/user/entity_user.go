package user

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

type UserDataInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
}

type UserServiceInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
}
