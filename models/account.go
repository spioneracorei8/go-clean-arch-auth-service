package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Account struct {
	Id                *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserId            *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4()" json:"user_id"`
	Username          string     `gorm:"type:text" json:"username"`
	PasswordPlainText string     `gorm:"-" json:"password"`
	PasswordBcrypt    string     `gorm:"type:text" json:"-"`
	WebAccess         string     `gorm:"type:web_access" json:"web_access"` // e.g., "APPLICATION", "MANAGEMENT"
	Status            string     `gorm:"type:account_status" json:"status"` // e.g., ACTIVE, INACTIVE

	CreatedBy string    `gorm:"type:text" json:"created_by"`
	UpdatedBy string    `gorm:"type:text" json:"updated_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Account) TableName() string {
	return "account"
}
