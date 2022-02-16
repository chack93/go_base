package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (m *Model) SetInit() {
	m.ID = uuid.New()
	m.CreatedAt = time.Now()
}

func (m *Model) SetUpdate() {
	m.UpdatedAt = time.Now()
}
