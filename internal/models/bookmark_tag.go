package models

import (
	"time"
)

type BookmarkTag struct {
	ID         int64     `json:"id" db:"id"`
	BookmarkID int64     `json:"bookmark_id" db:"bookmark_id"`
	TagID      int64     `json:"tag_id" db:"tag_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`

	Bookmark Bookmark `json:"bookmark" db:"bookmark" gorm:"foreignKey:BookmarkID"`
	Tag      Tag      `json:"tag" db:"tag" gorm:"foreignKey:TagID"`
}
