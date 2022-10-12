package params

type CreateUser struct {
	Nama      string
	Email     string
	Pekerjaan string
}

func NewCreateUser(nama string, email string, pekerjaan string) *CreateUser {
	return &CreateUser{
		Nama:      nama,
		Email:     email,
		Pekerjaan: pekerjaan,
	}
}
