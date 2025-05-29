package repository

import (
	"auth-service/models"
	"auth-service/services/register"

	"gorm.io/gorm"
)

type registerRepo struct {
	db *gorm.DB
}

func NewRegisterRepoImpl(db *gorm.DB) register.RegisterRepository {
	return &registerRepo{
		db: db,
	}
}

func (r *registerRepo) CreateAccount(account *models.Account) error {
	tx := r.db.Begin()

	if err := tx.Create(&account).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
