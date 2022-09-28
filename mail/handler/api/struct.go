package handlerAPI

import (
	"mime/multipart"
	"time"

	helperTime "github.com/Jiran03/mailku/helpers/time"
	"github.com/Jiran03/mailku/mail/domain"
)

type RequestJSON struct {
	Sender            string         `json:"sender" form:"sender" validate:"required"`
	ReferenceNumber   string         `json:"reference_number" form:"reference_number" validate:"required"`
	LetterDate        string         `json:"letter_date" form:"letter_date" validate:"required"`
	ReceivingAgencyID int            `json:"receiving_agency" form:"receiving_agency" validate:"required"`
	Receiver          string         `json:"receiver" form:"receiver" validate:"required"`
	DateOfReceipt     string         `json:"date_of_receipt" form:"date_of_receipt" validate:"required"`
	Receipt           multipart.File `form:"receipt"`
}

func toDomain(req RequestJSON) domain.Mail {
	return domain.Mail{
		Sender:            req.Sender,
		ReferenceNumber:   req.ReferenceNumber,
		LetterDate:        req.LetterDate,
		ReceivingAgencyID: req.ReceivingAgencyID,
		Receiver:          req.Receiver,
		DateOfReceipt:     req.DateOfReceipt,
		Receipt:           req.Receipt,
	}
}

type ResponseJSON struct {
	ID                int       `json:"id"`
	IDX               string    `json:"idx"`
	Sender            string    `json:"sender" form:"sender"`
	ReferenceNumber   string    `json:"reference_number" form:"reference_number"`
	LetterDate        string    `json:"letter_date" form:"letter_date"`
	ReceivingAgencyID int       `json:"receiving_agency" form:"receiving_agency"`
	Receiver          string    `json:"receiver" form:"receiver"`
	DateOfReceipt     string    `json:"date_of_receipt" form:"date_of_receipt"`
	ReceiptLink       string    `json:"receipt_link" form:"receipt_link"`
	CreatedAt         time.Time `json:"created_at" form:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Mail) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:                domain.ID,
		IDX:               domain.IDX,
		Sender:            domain.Sender,
		ReferenceNumber:   domain.ReferenceNumber,
		LetterDate:        domain.LetterDate,
		ReceivingAgencyID: domain.ReceivingAgencyID,
		Receiver:          domain.Receiver,
		DateOfReceipt:     domain.DateOfReceipt,
		ReceiptLink:       domain.ReceiptLink,
		CreatedAt:         tmCreatedAt,
		UpdatedAt:         tmUpdatedAt,
	}
}
