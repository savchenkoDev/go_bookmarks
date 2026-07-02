package repository

import (
	"bookmarks/internal/models"

	"gorm.io/gorm"
)

type BookmarkTagRepository struct {
	db *gorm.DB
}

func NewBookmarkTagRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

func (r *BookmarkRepository) DetachTagFromBookmark(userID int64, bookmarkTagID int64) error {
	err := r.db.Where("user_id = ? AND id = ?", userID, bookmarkTagID).Delete(&models.BookmarkTag{}).Error
	if err != nil {
		return err
	}
	return nil
}
