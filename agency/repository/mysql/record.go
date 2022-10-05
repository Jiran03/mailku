package repoMySQL

import (
	"github.com/Jiran03/mailku/agency/domain"
	mailRepo "github.com/Jiran03/mailku/mail/repository/mysql"
	"gorm.io/gorm"
)

type Agency struct {
	gorm.Model
	ID        int
	IDX       string
	Name      string
	CreatedAt string
	UpdatedAt string
	Mails     []mailRepo.Mail `gorm:"foreignKey:ReceivingAgencyID"`
}

func toDomain(rec Agency) domain.Agency {
	return domain.Agency{
		ID:        rec.ID,
		IDX:       rec.IDX,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Agency) Agency {
	return Agency{
		ID:        rec.ID,
		IDX:       rec.IDX,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
