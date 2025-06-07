package models

import (
	"time"

	"github.com/google/uuid"
	_"c02-project/docs"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string `json:"name"`
	Wave      int `gorm:"default:0" json:"wave"`
	CreatedAt time.Time `json:"created_at"`
}
