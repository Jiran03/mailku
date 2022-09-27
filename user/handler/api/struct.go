package handlerAPI

import (
	"time"

	helperTime "github.com/Jiran03/mailku/helpers/time"
	"github.com/Jiran03/mailku/user/domain"
)

type RequestJSON struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type RequestLoginJSON struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type Token struct {
	Token string `json:"token"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}
}

type ResponseJSON struct {
	ID        int       `json:"id"`
	IDX       string    `json:"idx"`
	Name      string    `json:"name" form:"name"`
	Username  string    `json:"username" form:"username"`
	Role      string    `json:"role" form:"role"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.User) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:        domain.ID,
		IDX:       domain.IDX,
		Name:      domain.Name,
		Username:  domain.Username,
		Role:      domain.Role,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
