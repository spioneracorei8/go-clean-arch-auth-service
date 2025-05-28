package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserConsent struct {
	UserId    *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"user_id"`
	ConsentId *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"consent_id"`
	Source    string     `gorm:"size:255" json:"source"` // e.g., "WEB_APPLICATION", "MOBILE_APPLICATION"
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

func (UserConsent) TableName() string {
	return "user_consent"
}
