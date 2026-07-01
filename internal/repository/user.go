package repository

import (
	"database/sql"
	"bookmarks/internal/user"
)

func GetUserByID(db *sql.DB, id int64) (user.User, error) {
	var u user.User
	err := db.QueryRow("SELECT id, email FROM users WHERE id = $1", id).Scan(&u.ID, &u.Email)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func GetUserByEmail(db *sql.DB, email string) (user.User, error) {
	var u user.User
	err := db.QueryRow("SELECT id, email, password_hash FROM users WHERE email = $1", email).Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}
