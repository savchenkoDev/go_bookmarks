package repository

import (
	"bookmarks/internal/models"

	"gorm.io/gorm"
)

type BookmarkTagRepository struct {
	db *gorm.DB
}

func NewBookmarkTagRepository(db *gorm.DB) *BookmarkTagRepository {
	return &BookmarkTagRepository{db: db}
}

func (r *BookmarkTagRepository) DetachTagFromBookmark(bookmarkTagID int64) error {
	err := r.db.Where("id = ?", bookmarkTagID).Delete(&models.BookmarkTag{}).Error
	if err != nil {
		return err
	}
	return nil
}
