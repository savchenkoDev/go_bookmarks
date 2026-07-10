package models

import (
	"time"
)

type BookmarkRequest struct {
	UserID      int64  `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	URL         string `json:"url" db:"url"`
	Description string `json:"description" db:"description"`
	IsFavorite  bool   `json:"is_favorite" db:"is_favorite"`
	IsArchived  bool   `json:"is_archived" db:"is_archived"`
}

type BookmarkUpdateRequest struct {
	Title       *string `json:"title"`
	URL         *string `json:"url"`
	Description *string `json:"description"`
	IsFavorite  *bool   `json:"is_favorite"`
	IsArchived  *bool   `json:"is_archived"`
}

type BookmarkResponse struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	URL         string    `json:"url" db:"url"`
	Description string    `json:"description" db:"description"`
	IsFavorite  bool      `json:"is_favorite" db:"is_favorite"`
	IsArchived  bool      `json:"is_archived" db:"is_archived"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	Tags []TagResponse `json:"tags"`
}

type Bookmark struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	URL         string    `json:"url" db:"url"`
	Description string    `json:"description" db:"description"`
	IsFavorite  bool      `json:"is_favorite" db:"is_favorite"`
	IsArchived  bool      `json:"is_archived" db:"is_archived"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	Tags []Tag `json:"tags" db:"tags" gorm:"many2many:bookmark_tags;"`
}

type BookmarkListParams struct {
	Page       int
	PerPage    int
	Sort       string
	Order      string
	IsFavorite *bool
	IsArchived *bool
	Tag        string
	Query      string
}

type PaginatedBookmarks struct {
	Data       []BookmarkResponse `json:"data"`
	Total      int64              `json:"total"`
	Page       int                `json:"page"`
	PerPage    int                `json:"per_page"`
	TotalPages int                `json:"total_pages"`
}

func (b *Bookmark) ToResponse() BookmarkResponse {
	tags := make([]TagResponse, len(b.Tags))
	for i, tag := range b.Tags {
		tags[i] = tag.ToResponse()
	}
	return BookmarkResponse{
		ID:          b.ID,
		Title:       b.Title,
		URL:         b.URL,
		Description: b.Description,
		IsFavorite:  b.IsFavorite,
		IsArchived:  b.IsArchived,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
		Tags:        tags,
	}
}
