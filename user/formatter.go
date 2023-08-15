package user

type Formatter struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func UserFormatter(user Users, token string) Formatter {
	userFormatter := Formatter{
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return userFormatter
}
