package handler

type UserResponRegister struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Notelp   string `json:"no_telp"`
	Respon   string `json:"respon"`
}

type ListCostumerRespon struct {
	Uuid          string `json:"uuid" form:"uuid"`
	UserName      string `json:"user_name" form:"user_name"`
	Status        string `json:"status" form:"status"`
	Role          string `json:"role" form:"role"`
	FullName      string `json:"full_name" form:"full_name"`
	TempatLahir   string `json:"tempat_lahir" form:"tempat_lahir"`
	Alamat        string `json:"alamat" form:"alamat"`
	Email         string `json:"email" form:"email"`
	Notelp        string `json:"notelp" form:"notelp"`
	NotelpKerabat string `json:"notelp_kerabat" form:"notelp_kerabat"`
	Ktp           string `json:"ktp" form:"ktp"`
	Pekerjaan     string `json:"pekerjaan" form:"pekerjaan"`
	FotoKtp       string `json:"foto_ktp" form:"foto_ktp"`
}

func CoreRespon(data ListCostumerRespon) ListCostumerRespon {
	return ListCostumerRespon{
		Uuid:          data.Uuid,
		UserName:      data.UserName,
		Status:        data.Status,
		Role:          data.Role,
		FullName:      data.FullName,
		TempatLahir:   data.TempatLahir,
		Alamat:        data.Alamat,
		Email:         data.Email,
		Notelp:        data.Notelp,
		NotelpKerabat: data.NotelpKerabat,
		Ktp:           data.Ktp,
		Pekerjaan:     data.Pekerjaan,
		FotoKtp:       data.FotoKtp,
	}
}

func CoreToList(data []ListCostumerRespon) []ListCostumerRespon {
	var result []ListCostumerRespon
	for _, v := range data {
		result = append(result, CoreRespon(v))
	}
	return result
}
