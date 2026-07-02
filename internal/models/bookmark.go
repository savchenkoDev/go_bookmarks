package models

import (
	"time"
)

type BookmarkRequest struct {
	UserID      int64 `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	URL         string `json:"url" db:"url"`
	Description string `json:"description" db:"description"`
	IsFavorite  bool `json:"is_favorite" db:"is_favorite"`
	IsArchived  bool `json:"is_archived" db:"is_archived"`
}

type BookmarkUpdateRequest struct {
	Title       *string `json:"title"`
	URL         *string `json:"url"`
	Description *string `json:"description"`
	IsFavorite  *bool   `json:"is_favorite"`
	IsArchived  *bool   `json:"is_archived"`
}

type BookmarkResponse struct {
	ID          int64 `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	URL         string `json:"url" db:"url"`
	Description string `json:"description" db:"description"`
	IsFavorite  bool `json:"is_favorite" db:"is_favorite"`
	IsArchived  bool `json:"is_archived" db:"is_archived"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	
	Tags []Tag `json:"tags" db:"tags"`
}

type Bookmark struct {
	ID          int64 `json:"id" db:"id"`
	UserID      int64 `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	URL         string `json:"url" db:"url"`
	Description string `json:"description" db:"description"`
	IsFavorite  bool `json:"is_favorite" db:"is_favorite"`
	IsArchived  bool `json:"is_archived" db:"is_archived"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	Tags []Tag `json:"tags" db:"tags" gorm:"many2many:bookmark_tags;"`
}

func (b *Bookmark) ToResponse() BookmarkResponse {
	return BookmarkResponse{
		ID: b.ID,
		Title: b.Title,
		URL: b.URL,
		Description: b.Description,
		IsFavorite: b.IsFavorite,
		IsArchived: b.IsArchived,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		Tags: b.Tags,
	}
}	