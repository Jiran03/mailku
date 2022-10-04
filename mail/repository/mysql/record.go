package repoMySQL

import (
	"github.com/Jiran03/mailku/mail/domain"
	"gorm.io/gorm"
)

type Mail struct {
	gorm.Model
	ID                int
	IDX               string
	Sender            string
	ReferenceNumber   string
	LetterDate        string
	ReceivingAgencyID int
	Receiver          string
	DateOfReceipt     string
	ReceiptLink       string
	CreatedAt         string
	UpdatedAt         string
}

func toDomain(rec Mail) domain.Mail {
	return domain.Mail{
		ID:                rec.ID,
		IDX:               rec.IDX,
		Sender:            rec.Sender,
		ReferenceNumber:   rec.ReferenceNumber,
		LetterDate:        rec.LetterDate,
		ReceivingAgencyID: rec.ReceivingAgencyID,
		Receiver:          rec.Receiver,
		DateOfReceipt:     rec.DateOfReceipt,
		ReceiptLink:       rec.ReceiptLink,
		CreatedAt:         rec.CreatedAt,
		UpdatedAt:         rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Mail) Mail {
	return Mail{
		ID:                rec.ID,
		IDX:               rec.IDX,
		Sender:            rec.Sender,
		ReferenceNumber:   rec.ReferenceNumber,
		LetterDate:        rec.LetterDate,
		ReceivingAgencyID: rec.ReceivingAgencyID,
		Receiver:          rec.Receiver,
		DateOfReceipt:     rec.DateOfReceipt,
		ReceiptLink:       rec.ReceiptLink,
		CreatedAt:         rec.CreatedAt,
		UpdatedAt:         rec.UpdatedAt,
	}
}
