package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
