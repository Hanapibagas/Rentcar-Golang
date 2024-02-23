package handler

type UserResponRegister struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Token      string `json:"token"`
}
