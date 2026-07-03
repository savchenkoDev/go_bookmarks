package models

import (
	"time"
)

type UserStats struct {
    Bookmarks int64 `json:"bookmarks" gorm:"column:total_bookmarks"`
    Favorites int64 `json:"favorites" gorm:"column:total_favorites"`
    Archived  int64 `json:"archived" gorm:"column:total_archived"`
    Tags      int64 `json:"tags" gorm:"column:total_tags"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty"`
}

type UserResponse struct {
	ID        int64 `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	ID        int64 `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	PasswordHash  string `json:"password_hash" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID: u.ID,
		Email: u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
