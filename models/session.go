package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// session used for just user that we want to login just one user for one account
type Session struct {
	Id        *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserId    *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4()" json:"user_id"`
	Source    string     `gorm:"size:255" json:"source"`     // e.g., "WEB_APPLICATION", "MOBILE_APPLICATION"
	Device    string     `gorm:"size:255" json:"device"`     // e.g., "IOS", "ANDROID", "WEB"
	TokenType string     `gorm:"size:255" json:"token_type"` // e.g., "ACCESS_TOKEN"
	CreatedAt time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp" json:"updated_at"`
}

func (Session) TableName() string {
	return "session"
}

func (s *Session) GenUUID() {
	id, _ := uuid.NewV4()
	s.Id = &id
}
