package service

import (
	"github.com/Jiran03/mailku/agency/domain"
	authMiddleware "github.com/Jiran03/mailku/auth/middlewares"
	timeHelper "github.com/Jiran03/mailku/helpers/time"
	"github.com/google/uuid"
)

type agencyService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// UpdateData implements domain.Service
func (as agencyService) UpdateData(id string, domain domain.Agency) (agencyObj domain.Agency, err error) {
	if agencyObj, err = as.GetByID(id); err != nil {
		return agencyObj, err
	}

	domain.ID = agencyObj.ID
	domain.IDX = id
	domain.CreatedAt = agencyObj.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	if agencyObj, err = as.repository.Update(domain); err != nil {
		return agencyObj, err
	}

	return agencyObj, nil
}

// GetByID implements domain.Service
func (as agencyService) GetByID(id string) (agencyObj domain.Agency, err error) {
	if agencyObj, err = as.repository.GetByID(id); err != nil {
		return agencyObj, err
	}

	return agencyObj, nil
}

func (as agencyService) InsertData(domain domain.Agency) (agencyObj domain.Agency, err error) {
	domain.IDX = uuid.New().String()
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	if agencyObj, err = as.repository.Create(domain); err != nil {
		return agencyObj, err
	}

	return agencyObj, nil
}

// GetAllData implements domain.Service
func (as agencyService) GetAllData() (agencyObj []domain.Agency, err error) {
	if agencyObj, err = as.repository.Get(); err != nil {
		return agencyObj, err
	}

	return agencyObj, nil
}

// DeleteData implements domain.Service
func (as agencyService) DeleteData(id string) (err error) {
	if errResp := as.repository.Delete(id); errResp != nil {
		return errResp
	}

	return nil
}

func NewAgencyService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return agencyService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
