package handlerAPI

import (
	"time"

	"github.com/Jiran03/mailku/agency/domain"
	helperTime "github.com/Jiran03/mailku/helpers/time"
)

type RequestJSON struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func toDomain(req RequestJSON) domain.Agency {
	return domain.Agency{
		Name: req.Name,
	}
}

type ResponseJSON struct {
	ID        int       `json:"id"`
	IDX       string    `json:"idx"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Agency) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:        domain.ID,
		IDX:       domain.IDX,
		Name:      domain.Name,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
