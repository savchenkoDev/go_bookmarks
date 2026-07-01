package repository

import (
	"gorm.io/gorm"

	"bookmarks/internal/tag"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Create(userID int64, name string) (tag.Tag, error) {
	var t tag.Tag
	err := r.db.Create(&tag.Tag{UserID: userID, Name: name}).Error
	if err != nil {
		return tag.Tag{}, err
	}
	return t, nil
}