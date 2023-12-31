package user

type (
	InputRegistUser struct {
		Name       string `json:"name" binding:"required"`
		Occupation string `json:"occupation" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
	}

	InputLoginUser struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	CheckedEmailInput struct {
		Email string `json:"email" binding:"required,email"`
	}
)
