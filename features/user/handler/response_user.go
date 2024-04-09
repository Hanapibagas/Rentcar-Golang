package handler

type UserResponRegister struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Notelp   string `json:"no_telp"`
	Token    string `json:"token"`
	Respon   string `json:"respon"`
}
