package user

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
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

func (ur *UserRequest) Create(db *sql.DB) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	var u User
	err = db.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, email, created_at, updated_at", ur.Email, passwordHash).Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	
	if err != nil {
		return User{}, err
	}
	return u, nil
}