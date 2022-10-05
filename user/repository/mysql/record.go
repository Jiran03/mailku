package repoMySQL

import (
	"github.com/Jiran03/mailku/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	IDX       string
	Name      string
	Username  string
	Password  string
	Role      string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec User) domain.User {
	return domain.User{
		ID:        rec.ID,
		IDX:       rec.IDX,
		Name:      rec.Name,
		Username:  rec.Username,
		Password:  rec.Password,
		Role:      rec.Role,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.User) User {
	return User{
		ID:        rec.ID,
		IDX:       rec.IDX,
		Name:      rec.Name,
		Username:  rec.Username,
		Password:  rec.Password,
		Role:      rec.Role,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
