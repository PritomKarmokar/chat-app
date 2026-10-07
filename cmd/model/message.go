package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	SenderID   uuid.UUID `gorm:"column:sender_id; not null" json:"sender_id"`
	ReceiverID uuid.UUID `gorm:"column:receiver_id; not null" json:"receiver_id"`
	Content    string    `gorm:"column:content; type:text; not null" json:"content"`
	CreatedAt  time.Time `gorm:"column:created_at; not null" json:"created_at"`
}

func (m *Message) BeforeCreate(db *gorm.DB) error {
	if m.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		m.ID = id
	}
	return nil
}

func (m *Message) TableName() string {
	return "messages"
}
