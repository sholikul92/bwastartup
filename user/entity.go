package user

import "time"

type (
	Users struct {
		ID             int
		Name           string
		Email          string
		Occupation     string
		PasswordHash   string
		AvatarFileName string
		Role           string
		Token          string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)
