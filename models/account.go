package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id                *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserId            *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4()" json:"user_id"`
	Username          string     `gorm:"size:255" json:"username"`
	PasswordPlainText string     `gorm:"-" json:"password"`
	PasswordBcrypt    string     `gorm:"type:text;column:password" json:"-"`
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

func (a *Account) GenUUID() {
	id, _ := uuid.NewV4()
	a.Id = &id
}

func (a *Account) BcryptPwd() {
	bcryptPwd, err := bcrypt.GenerateFromPassword([]byte(a.PasswordPlainText), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal("Error generating bcrypt password:", err.Error())
		return
	}
	a.PasswordBcrypt = string(bcryptPwd)
}
