package bookmark

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
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
}

func (br *BookmarkRequest) Create(db *sql.DB, userID int64) (Bookmark, error) {
	var b Bookmark
	query := "INSERT INTO bookmarks (user_id, title, url, description, is_favorite, is_archived) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, title, url, description, is_favorite, is_archived, created_at, updated_at"
	err := db.QueryRow(query, userID, br.Title, br.URL, br.Description, br.IsFavorite, br.IsArchived).Scan(&b.ID, &b.UserID, &b.Title, &b.URL, &b.Description, &b.IsFavorite, &b.IsArchived, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return Bookmark{}, err
	}
	return b, nil
}

func (br *BookmarkUpdateRequest) Update(db *sql.DB, id int64, userID int64) (Bookmark, error) {
	setParts := []string{}
	args := []any{}
	n := 1

	if br.Title != nil {
		setParts = append(setParts, "title = $"+strconv.Itoa(n))
		args = append(args, *br.Title)
		n++
	}
	if br.URL != nil {
		setParts = append(setParts, "url = $"+strconv.Itoa(n))
		args = append(args, *br.URL)
		n++
	}
	if br.Description != nil {
		setParts = append(setParts, "description = $"+strconv.Itoa(n))
		args = append(args, *br.Description)
		n++
	}
	if br.IsFavorite != nil {
		setParts = append(setParts, "is_favorite = $"+strconv.Itoa(n))
		args = append(args, *br.IsFavorite)
		n++
	}
	if br.IsArchived != nil {
		setParts = append(setParts, "is_archived = $"+strconv.Itoa(n))
		args = append(args, *br.IsArchived)
		n++
	}

	if len(setParts) == 0 {
		return Bookmark{}, errors.New("nothing to update")
	}

	setParts = append(setParts, "updated_at = NOW()")

	idPlaceholder := strconv.Itoa(n)
	userPlaceholder := strconv.Itoa(n + 1)
	args = append(args, id, userID)

	query := fmt.Sprintf(
		`UPDATE bookmarks
		 SET %s
		 WHERE id = $%s AND user_id = $%s
		 RETURNING id, user_id, title, url, description, is_favorite, is_archived, created_at, updated_at`,
		strings.Join(setParts, ", "),
		idPlaceholder,
		userPlaceholder,
	)

	var b Bookmark
	err := db.QueryRow(query, args...).Scan(
		&b.ID, &b.UserID, &b.Title, &b.URL, &b.Description,
		&b.IsFavorite, &b.IsArchived, &b.CreatedAt, &b.UpdatedAt,
	)
	if err != nil {
		return Bookmark{}, err
	}

	return b, nil
}

func (b *Bookmark) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM bookmarks WHERE id = $1", b.ID)
	return err
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
	}
}