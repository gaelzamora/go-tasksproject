package domain

import (
	"time"
)

type Task struct {
    ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
    CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
    Name      string     `json:"name"`
    Content   string     `json:"content"`
    UserID    uint       `json:"user_id"`
}