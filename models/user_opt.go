package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserOTP struct {
	Id          *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserId      *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4()" json:"user_id"`
	Email       string     `gorm:"size:255" json:"email"`              // Email address for OTP
	Source      string     `gorm:"size:255" json:"source"`             // e.g., "APPLICATION" = for user, "MANAGEMENT for admin"
	OTP         string     `gorm:"size:255" json:"otp"`                // One-time password
	RefKey      string     `gorm:"size:255" json:"ref_key"`            // Reference key for OTP
	IsUsed      bool       `gorm:"default:false" json:"is_used"`       // Indicates if the OTP has been used
	IsValidated bool       `gorm:"default:false" json:"is_validated"`  // Indicates if the password has been changed
	ExpiredDate time.Time  `gorm:"type:timestamp" json:"expired_date"` // Expiration date for the OTP
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`   // Creation timestamp
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`   // Update timestamp
}

func (UserOTP) TableName() string {
	return "user_otp"
}
func (u *UserOTP) GenUUID() {
	id, _ := uuid.NewV4()
	u.Id = &id
}
