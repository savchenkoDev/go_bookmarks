package repository

import (
	"bookmarks/internal/models"

	"gorm.io/gorm"
)

type BookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

func (r *BookmarkRepository) GetBookmarkByIDAndUserID(id int64, userID int64) (models.Bookmark, error) {
	var b models.Bookmark
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	return b, nil
}

func (r *BookmarkRepository) GetBookmarksByUserID(userID int64) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	err := r.db.Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}
	return bookmarks, nil
}

func (r *BookmarkRepository) Create(userID int64, br models.BookmarkRequest) (models.Bookmark, error) {
	var b models.Bookmark
	b = models.Bookmark{
		UserID: userID,
		Title: br.Title,
		URL: br.URL,
		Description: br.Description,	
		IsFavorite: br.IsFavorite,
		IsArchived: br.IsArchived,
	}
	err := r.db.Create(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	return b, nil
}

func (r *BookmarkRepository) Delete(userID int64, id int64) error {
	err := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Bookmark{}).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&models.Bookmark{}, id).Error
}

func (r *BookmarkRepository) Update(userID int64, id int64, br models.BookmarkUpdateRequest) (models.Bookmark, error) {
	var b models.Bookmark
	err := r.db.Where("user_id = ? AND id = ?", userID, id).Updates(br).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	err = r.db.Save(&b).Error
	if err != nil {
		return models.Bookmark{}, err
	}
	return b, nil
}

func (r *BookmarkRepository) AttachTagToBookmark(userID int64, bookmarkID int64, tagID int64) (models.BookmarkTag, error ){
	bt := models.BookmarkTag{
		BookmarkID: bookmarkID,
		TagID: tagID,
	}
	err := r.db.Create(&bt).Error
	if err != nil {
		return models.BookmarkTag{}, err
	}

	err = r.db.
			Preload("Bookmark").
			Preload("Tag").
			First(&bt, bt.ID).Error
	if err != nil {
		return models.BookmarkTag{}, err
	}
	return bt, nil
}
