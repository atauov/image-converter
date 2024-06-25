package models

import "time"

type Image struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UserID      int       `gorm:"user_id" json:"user_id"`
	Filename    string    `gorm:"filename" json:"filename"`
	OriginalUrl string    `gorm:"original_url" json:"original_url"`
	SizeLarge   string    `gorm:"size_large" json:"size_large"`
	SizeMedium  string    `gorm:"size_medium" json:"size_medium"`
	SizeSmall   string    `gorm:"size_small" json:"size_small"`
	IsDone      bool      `gorm:"is_done" json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
