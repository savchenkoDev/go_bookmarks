package models

import "time"

type Tag struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type TagRequest struct {
	UserID int64  `json:"user_id" db:"user_id"`
	Name   string `json:"name" binding:"required,min=4,max=20"`
}

type TagUpdateRequest struct {
	Name *string `json:"name" binding:"omitempty"`
}

type TagResponse struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" binding:"required,min=4,max=20"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (t *Tag) ToResponse() TagResponse {
	return TagResponse{
		ID:   t.ID,
		Name: t.Name,
	}
}
