package domain

import "mime/multipart"

type Mail struct {
	ID                int
	IDX               string
	Sender            string
	ReferenceNumber   string // nomor surat
	LetterDate        string
	ReceivingAgencyID int // instansi penerima
	Receiver          string
	DateOfReceipt     string // tanggal terima
	Receipt           multipart.File
	ReceiptLink       string
	CreatedAt         string
	UpdatedAt         string
}
