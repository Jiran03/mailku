package service

import (
	"bytes"
	"io"

	authMiddleware "github.com/Jiran03/mailku/auth/middlewares"
	uploadImageHelper "github.com/Jiran03/mailku/helpers/azure"
	timeHelper "github.com/Jiran03/mailku/helpers/time"
	"github.com/Jiran03/mailku/mail/domain"
	"github.com/google/uuid"
)

type mailService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// UpdateData implements domain.Service
func (ms mailService) UpdateData(id string, domain domain.Mail) (mailObj domain.Mail, err error) {
	if mailObj, err = ms.GetByID(id); err != nil {
		return mailObj, err
	}

	domain.ID = mailObj.ID
	domain.IDX = id
	if domain.Receipt != nil {
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, domain.Receipt); err != nil {
			return mailObj, err
		}

		data := buf.Bytes()
		if domain.ReceiptLink, err = uploadImageHelper.UploadBytesToBlob(data); err != nil {
			return mailObj, err
		}

	} else {
		domain.ReceiptLink = mailObj.ReceiptLink
	}

	domain.CreatedAt = mailObj.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	if mailObj, err = ms.repository.Update(domain); err != nil {
		return mailObj, err
	}

	return mailObj, nil
}

// GetByID implements domain.Service
func (ms mailService) GetByID(id string) (mailObj domain.Mail, err error) {
	if mailObj, err = ms.repository.GetByID(id); err != nil {
		return mailObj, err
	}

	return mailObj, nil
}

func (ms mailService) InsertData(domain domain.Mail) (mailObj domain.Mail, err error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, domain.Receipt); err != nil {
		return mailObj, err
	}

	domain.IDX = uuid.New().String()
	data := buf.Bytes()
	if domain.ReceiptLink, err = uploadImageHelper.UploadBytesToBlob(data); err != nil {
		return mailObj, err
	}

	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	if mailObj, err = ms.repository.Create(domain); err != nil {
		return mailObj, err
	}

	return mailObj, nil
}

// GetAllData implements domain.Service
func (ms mailService) GetAllData() (mailObj []domain.Mail, err error) {
	if mailObj, err = ms.repository.Get(); err != nil {
		return mailObj, err
	}

	return mailObj, nil
}

// DeleteData implements domain.Service
func (ms mailService) DeleteData(id string) (err error) {
	if errResp := ms.repository.Delete(id); errResp != nil {
		return errResp
	}

	return nil
}

func NewMailService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return mailService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
