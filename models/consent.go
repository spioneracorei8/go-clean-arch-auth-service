package models

import "github.com/gofrs/uuid"

type Consent struct {
	Id      *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Source  string     `gorm:"size:255" json:"source"` // e.g., "WEB_APPLICATION", "MOBILE_APPLICATION"
	Title   string     `gorm:"size:255" json:"title"`
	Content string     `gorm:"type:text" json:"content"`
}

func (Consent) TableName() string {
	return "consent"
}

func (c *Consent) GenUUID() {
	id, _ := uuid.NewV4()
	c.Id = &id
}
