package repoMySQL

import (
	"github.com/Jiran03/mailku/agency/domain"
	"gorm.io/gorm"
)

type agencyRepository struct {
	DB *gorm.DB
}

// Update implements domain.Repository
func (ar agencyRepository) Update(domain domain.Agency) (agencyObj domain.Agency, err error) {
	var newRecord Agency
	rec := fromDomain(domain)
	if err = ar.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":         rec.ID,
		"id_x":       rec.IDX,
		"name":       rec.Name,
		"created_at": rec.CreatedAt,
		"updated_at": rec.UpdatedAt,
	}).Error; err != nil {
		return agencyObj, err
	}

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (ar agencyRepository) GetByID(id string) (domain domain.Agency, err error) {
	var record Agency
	err = ar.DB.Where("id_x = ?", id).First(&record).Error
	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

// Get implements domain.Repository
func (ar agencyRepository) Get() (agencyObj []domain.Agency, err error) {
	var records []Agency
	err = ar.DB.Find(&records).Error
	if err != nil {
		return agencyObj, err
	}

	for _, value := range records {
		agencyObj = append(agencyObj, toDomain(value))
	}

	return agencyObj, nil
}

// Create implements domain.Repository
func (ar agencyRepository) Create(domain domain.Agency) (agencyObj domain.Agency, err error) {
	// var recordDetail AgencyDetail
	record := fromDomain(domain)
	err = ar.DB.Create(&record).Error
	if err != nil {
		return agencyObj, err
	}

	agencyObj = toDomain(record)
	return agencyObj, nil
}

// Delete implements domain.Repository
func (ar agencyRepository) Delete(id string) (err error) {
	var record Agency
	if err = ar.DB.Where("id_x = ?", id).Delete(&record).Error; err != nil {
		return err
	}

	return nil
}

func NewAgencyRepository(db *gorm.DB) domain.Repository {
	return agencyRepository{
		DB: db,
	}
}
