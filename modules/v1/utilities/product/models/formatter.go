package models

type UserFormatter struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    user.Id_users,
		Nama:  user.Nama,
		Email: user.Email,
		Token: token,
	}

	return formatter
}
