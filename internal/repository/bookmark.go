package repository

import (
	"database/sql"
	"bookmarks/internal/bookmark"
)

func GetBookmarkByIDAndUserID(db *sql.DB, id int64, userID int64) (bookmark.Bookmark, error) {
	var b bookmark.Bookmark
	err := db.QueryRow("SELECT id, title, url, description, is_favorite, is_archived, created_at, updated_at FROM bookmarks WHERE id = $1 AND user_id = $2", id, userID).Scan(&b.ID, &b.Title, &b.URL, &b.Description, &b.IsFavorite, &b.IsArchived, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return bookmark.Bookmark{}, err
	}
	return b, nil
}

func GetBookmarksByUserID(db *sql.DB, userID int64) ([]bookmark.Bookmark, error) {
	rows, err := db.Query("SELECT id, title, url, description, is_favorite, is_archived, created_at, updated_at FROM bookmarks WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []bookmark.Bookmark
	for rows.Next() {
		var b bookmark.Bookmark
		err = rows.Scan(&b.ID, &b.Title, &b.URL, &b.Description, &b.IsFavorite, &b.IsArchived, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, b)
	}
	return bookmarks, nil
}
