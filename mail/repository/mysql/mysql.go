package repoMySQL

import (
	"github.com/Jiran03/mailku/mail/domain"
	"gorm.io/gorm"
)

type mailRepository struct {
	DB *gorm.DB
}

// Update implements domain.Repository
func (mr mailRepository) Update(domain domain.Mail) (mailObj domain.Mail, err error) {
	var newRecord Mail
	rec := fromDomain(domain)
	if err = mr.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":                  rec.ID,
		"id_x":                rec.IDX,
		"sender":              rec.Sender,
		"reference_number":    rec.ReferenceNumber,
		"letter_date":         rec.LetterDate,
		"receiving_agency_id": rec.ReceivingAgencyID,
		"receiver":            rec.Receiver,
		"receipt_link":        rec.ReceiptLink,
		"created_at":          rec.CreatedAt,
		"updated_at":          rec.UpdatedAt,
	}).Error; err != nil {
		return mailObj, err
	}

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (mr mailRepository) GetByID(id string) (domain domain.Mail, err error) {
	var record Mail
	err = mr.DB.Where("id_x = ?", id).First(&record).Error

	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

// Get implements domain.Repository
func (mr mailRepository) Get() (mailObj []domain.Mail, err error) {
	var records []Mail
	err = mr.DB.Find(&records).Error
	if err != nil {
		return mailObj, err
	}

	for _, value := range records {
		mailObj = append(mailObj, toDomain(value))
	}

	return mailObj, nil
}

// Create implements domain.Repository
func (mr mailRepository) Create(domain domain.Mail) (mailObj domain.Mail, err error) {
	// var recordDetail MailDetail
	record := fromDomain(domain)
	err = mr.DB.Create(&record).Error
	if err != nil {
		return mailObj, err
	}

	mailObj = toDomain(record)
	return mailObj, nil
}

// Delete implements domain.Repository
func (mr mailRepository) Delete(id string) (err error) {
	var record Mail
	err = mr.DB.Delete(&record, id).Error
	if err != nil {
		return err
	}

	return nil
}

func NewMailRepository(db *gorm.DB) domain.Repository {
	return mailRepository{
		DB: db,
	}
}
