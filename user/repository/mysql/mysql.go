package repoMySQL

import (
	"errors"

	"github.com/Jiran03/mailku/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// Update implements domain.Repository
func (ur userRepository) Update(domain domain.User) (userObj domain.User, err error) {
	var newRecord User
	rec := fromDomain(domain)
	if err = ur.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":         rec.ID,
		"id_x":       rec.IDX,
		"name":       rec.Name,
		"username":   rec.Username,
		"password":   rec.Password,
		"role":       rec.Role,
		"created_at": rec.CreatedAt,
		"updated_at": rec.UpdatedAt,
	}).Error; err != nil {
		return userObj, err
	}

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (ur userRepository) GetByID(id string) (domain domain.User, err error) {
	var record User
	err = ur.DB.Where("id_x = ?", id).First(&record).Error

	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

// Get implements domain.Repository
func (ur userRepository) Get() (userObj []domain.User, err error) {
	var records []User
	err = ur.DB.Find(&records).Error
	if err != nil {
		return userObj, err
	}

	for _, value := range records {
		userObj = append(userObj, toDomain(value))
	}

	return userObj, nil
}

// GetByUsername implements domain.Repository
func (ur userRepository) GetByUsername(username string) (userObj domain.User, err error) {
	var record User
	err = ur.DB.Where("username = ?", username).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userObj, err
	}

	return toDomain(record), nil
}

// Create implements domain.Repository
func (ur userRepository) Create(domain domain.User) (userObj domain.User, err error) {
	// var recordDetail UserDetail
	record := fromDomain(domain)
	err = ur.DB.Create(&record).Error
	if err != nil {
		return userObj, err
	}

	userObj = toDomain(record)
	return userObj, nil
}

// Delete implements domain.Repository
func (ur userRepository) Delete(id int) (err error) {
	var record User
	err = ur.DB.Delete(&record, id).Error
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepository{
		DB: db,
	}
}
